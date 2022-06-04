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

package topic_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/bangumi/server/internal/dal/dao"
	"github.com/bangumi/server/internal/dal/query"
	"github.com/bangumi/server/internal/domain"
	"github.com/bangumi/server/internal/model"
	"github.com/bangumi/server/internal/test"
	"github.com/bangumi/server/internal/topic"
)

func getRepo(t *testing.T) domain.TopicRepo {
	t.Helper()
	repo, err := topic.NewMysqlRepo(query.Use(test.GetGorm(t)), zap.NewNop())
	require.NoError(t, err)

	return repo
}

func TestGet(t *testing.T) {
	test.RequireEnv(t, test.EnvMysql)
	t.Parallel()

	repo := getRepo(t)

	s, err := repo.Get(context.Background(), domain.TopicTypeSubject, 1, 0, 0)
	require.NoError(t, err)

	require.Equal(t, model.TopicIDType(1), s.ID)
}

func TestMysqlRepo_GetTopics(t *testing.T) {
	test.RequireEnv(t, test.EnvMysql)
	t.Parallel()

	repo := getRepo(t)

	_, err := repo.GetTopics(context.Background(), domain.TopicTypeSubject, 2)
	require.NoError(t, err)
}

func TestMysqlRepo_ConvertDao(t *testing.T) {
	t.Parallel()

	p, err := topic.ConvertDao(&dao.SubjectTopic{
		ID:        10,
		SubjectID: 20,
	})
	require.NoError(t, err)
	require.Equal(t, p.ID, model.TopicIDType(10))
	require.Equal(t, p.ObjectID, uint32(20))
}
