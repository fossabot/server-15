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

package dao

import (
	"time"

	"github.com/bangumi/server/internal/model"
)

func (c *GroupTopicComment) CreatorID() model.UserID {
	return model.UserID(c.UID)
}

func (c *GroupTopicComment) IsSubComment() bool {
	return c.Related == 0
}

func (c *GroupTopicComment) CommentID() model.CommentID {
	return c.ID
}

func (c *GroupTopicComment) CreateAt() time.Time {
	return time.Unix(int64(c.CreatedTime), 0)
}

func (c *GroupTopicComment) GetState() uint8 {
	return c.statStub()
}

func (c *GroupTopicComment) RelatedTo() model.CommentID {
	return c.Related
}

func (c *GroupTopicComment) GetID() model.CommentID {
	return c.ID
}

func (c *GroupTopicComment) GetContent() string {
	return c.Content
}

func (c *GroupTopicComment) GetMentionedID() model.UserID {
	return model.UserID(c.MentionedID)
}

func (c *GroupTopicComment) statStub() uint8 {
	return c.State
}
