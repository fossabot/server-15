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

package handler_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/bangumi/server/internal/mocks"
	"github.com/bangumi/server/internal/model"
	"github.com/bangumi/server/internal/pkg/test"
	"github.com/bangumi/server/internal/web/res"
)

func TestHandler_GetPerson_HappyPath(t *testing.T) {
	t.Parallel()
	m := mocks.NewPersonRepo(t)
	m.EXPECT().Get(mock.Anything, model.PersonID(7)).Return(model.Person{ID: 7}, nil)

	app := test.GetWebApp(t, test.Mock{PersonRepo: m})

	resp := test.New(t).Get("/v0/persons/7").Execute(app)

	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestHandler_GetPerson_Redirect(t *testing.T) {
	t.Parallel()
	m := mocks.NewPersonRepo(t)
	m.EXPECT().Get(mock.Anything, model.PersonID(7)).Return(model.Person{ID: 7}, nil)

	app := test.GetWebApp(t, test.Mock{PersonRepo: m})

	resp := test.New(t).Get("/v0/persons/7").Execute(app)

	require.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestHandler_GetPerson_Redirect_cached(t *testing.T) {
	t.Parallel()
	c := mocks.NewCache(t)
	c.EXPECT().Get(mock.Anything, mock.Anything, mock.Anything).Run(func(_ context.Context, _ string, value any) {
		reflect.ValueOf(value).Elem().Set(reflect.ValueOf(res.PersonV0{Redirect: 8}))
	}).Return(true, nil)

	app := test.GetWebApp(t, test.Mock{Cache: c})
	resp := test.New(t).Get("/v0/persons/7").Execute(app)

	require.Equal(t, http.StatusFound, resp.StatusCode)
	require.Equal(t, "/v0/persons/8", resp.Header.Get("Location"))
}

func TestHandler_GetPersonImage_302(t *testing.T) {
	t.Parallel()
	m := mocks.NewPersonRepo(t)
	m.EXPECT().Get(mock.Anything, model.PersonID(1)).Return(model.Person{ID: 1, Image: "temp"}, nil)
	m.EXPECT().Get(mock.Anything, model.PersonID(3)).Return(model.Person{ID: 3}, nil)

	app := test.GetWebApp(t, test.Mock{PersonRepo: m})

	for _, imageType := range []string{"small", "grid", "large", "medium"} {
		t.Run(imageType, func(t *testing.T) {
			t.Parallel()

			resp := test.New(t).Get("/v0/persons/1/image?type=" + imageType).Execute(app)
			require.Equal(t, http.StatusFound, resp.StatusCode, resp.BodyString())
			expected, _ := res.PersonImage("temp").Select(imageType)
			require.Equal(t, expected, resp.Header.Get("Location"), "expect redirect to image url")

			// should redirect to default image
			resp = test.New(t).Get("/v0/persons/3/image?type=" + imageType).Execute(app)
			require.Equal(t, http.StatusFound, resp.StatusCode, resp.BodyString())
			require.Equal(t, res.DefaultImageURL, resp.Header.Get("Location"), "should redirect to default image")
		})
	}
}

func TestHandler_GetPersonImage_400(t *testing.T) {
	t.Parallel()
	m := mocks.NewPersonRepo(t)
	m.EXPECT().Get(mock.Anything, mock.Anything).Return(model.Person{Image: "temp"}, nil)

	app := test.GetWebApp(t, test.Mock{PersonRepo: m})

	resp := test.New(t).Get("/v0/persons/1/image").Execute(app)
	require.Equal(t, http.StatusBadRequest, resp.StatusCode, resp.BodyString())
}
