// SPDX-License-Identifier: AGPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>

package handler

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"github.com/bangumi/server/internal/domain"
	"github.com/bangumi/server/internal/errgo"
	"github.com/bangumi/server/internal/model"
	"github.com/bangumi/server/internal/web/res"
	"github.com/bangumi/server/internal/web/util"
)

func (h Handler) listComments(c *fiber.Ctx, commentType domain.CommentType, id uint32) (*res.Paged, error) {
	page, err := getPageQuery(c, defaultPageLimit, defaultMaxPageLimit)
	if err != nil {
		return nil, c.Status(http.StatusNotFound).JSON(res.Error{
			Title:   "Not Found",
			Details: util.DetailFromRequest(c),
		})
	}

	count, err := h.m.Count(c.Context(), commentType, id)
	if err != nil {
		return nil, errgo.Wrap(err, "repo.comments.Count")
	}

	if count == 0 {
		return nil, c.JSON(res.Paged{Data: []res.Comment{}, Total: count, Limit: page.Limit, Offset: page.Offset})
	}

	if err = page.check(count); err != nil {
		return nil, err
	}

	comments, err := h.m.List(c.Context(), commentType, id, page.Limit, page.Offset)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, c.Status(http.StatusNotFound).JSON(res.Error{
				Title:   "Not Found",
				Details: util.DetailFromRequest(c),
			})
		}
		return nil, errgo.Wrap(err, "Comment.GetCommentsByMentionedID")
	}

	uids := make([]model.UIDType, 0)
	commentMap := make(map[model.CommentIDType]model.Comment)
	for _, v := range comments {
		commentMap[v.ID] = v
	}

	extIDs := make([]model.CommentIDType, 0)
	for _, v := range comments {
		if _, ok := commentMap[v.Related]; !ok && v.Related != 0 {
			extIDs = append(extIDs, v.Related)
		}
	}

	if len(extIDs) != 0 {
		relatedComments, e := h.m.GetByIDs(c.Context(), commentType, extIDs...)
		if e != nil {
			return nil, errgo.Wrap(e, "repo.comments.GetByIDs")
		}
		for _, v := range relatedComments {
			commentMap[v.ID] = v
		}
	}

	for _, v := range commentMap {
		uids = append(uids, v.UID)
	}

	userMap, err := h.u.GetByIDs(c.Context(), dedupeUIDs(uids...)...)
	if err != nil {
		return nil, errgo.Wrap(err, "user.GetByIDs")
	}

	return &res.Paged{
		Total:  count,
		Limit:  page.Limit,
		Offset: page.Offset,
		Data:   convertModelComments(comments, commentMap, userMap),
	}, nil
}

func convertModelComments(
	comments []model.Comment, cm map[model.CommentIDType]model.Comment, userMap map[uint32]model.User,
) []res.Comment {
	result := make([]res.Comment, len(comments))
	for k, v := range comments {
		result[k] = res.Comment{
			ID:        v.ID,
			Text:      v.Content,
			CreatedAt: v.CreatedAt,
			Creator:   convertModelUser(userMap[v.UID]),
		}
		if related, ok := cm[v.Related]; ok {
			result[k].ReplyTo = &res.Comment{
				CreatedAt: related.CreatedAt,
				Creator:   convertModelUser(userMap[related.UID]),
				Text:      related.Content,
				ID:        related.ID,
			}
		}
	}
	return result
}
