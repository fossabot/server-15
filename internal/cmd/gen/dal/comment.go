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

// generate methods for dal comments dao.
package main

import (
	_ "embed"
	"io/ioutil"
	"path/filepath"
	"strings"
)

//go:embed template/comments.go
var template string

func main() {
	for _, t := range []struct {
		Value        string
		Name         string
		GenStateStub bool
	}{
		{
			Value:        "PersonComment",
			Name:         "comment_person",
			GenStateStub: true,
		},
		{
			Value:        "CharacterComment",
			Name:         "comment_characters",
			GenStateStub: true,
		},
		{
			Value: "GroupTopicComment",
			Name:  "comment_group",
		},
		{
			Value: "SubjectTopicComment",
			Name:  "comment_subject",
		},
		{
			Value:        "EpisodeComment",
			Name:         "comment_episode",
			GenStateStub: true,
		},
		{
			Value:        "IndexComment",
			Name:         "comment_index",
			GenStateStub: true,
		},
	} {
		name := t.Value

		content := template

		if t.GenStateStub {
			content += `
func (c *TypeComment) statStub() uint8 {
	return 0
}
`
		} else {
			content += `
func (c *TypeComment) statStub() uint8 {
	return c.State
}
`
		}

		content = strings.ReplaceAll(content, "TypeComment", name)

		fileName := t.Name + ".go"
		err := ioutil.WriteFile(filepath.Join("./internal/dal/dao/", fileName), []byte(content), 0600) //nolint:gomnd
		if err != nil {
			panic(err)
		}
	}
}
