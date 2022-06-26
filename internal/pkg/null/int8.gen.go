// Code generated by ./internal/cmd/gen/null/main.go. DO NOT EDIT.

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

package null

import (
	"bytes"

	"github.com/goccy/go-json"
)

var _ json.Unmarshaler = (*Int8)(nil)

// Int8 is a nullable type.
type Int8 struct {
	Int8 int8
	Set  bool // if json object has this field
	Null bool // if json field's value is `null`
}

// NewInt8 creates a new int8.
func NewInt8(t int8) Int8 {
	return Int8{
		Null: false,
		Int8: t,
		Set:  true,
	}
}

func (t Int8) HasValue() bool {
	return t.Set && !t.Null
}

func (t Int8) Ptr() *int8 {
	if t.Set && !t.Null {
		return &t.Int8
	}

	return nil
}

func (t Int8) Interface() interface{} {
	if t.Set && !t.Null {
		return &t.Int8
	}

	return nil
}

// Default return default value its value is Null or not Set.
func (t Int8) Default(v int8) int8 {
	if t.Null && t.Set {
		return t.Int8
	}

	return v
}

// UnmarshalJSON implements json.Unmarshaler.
func (t *Int8) UnmarshalJSON(data []byte) error {
	t.Set = true

	if bytes.Equal(data, nilBytes) {
		t.Null = true
		return nil
	}

	if err := json.UnmarshalNoEscape(data, &t.Int8); err != nil {
		return err //nolint:wrapcheck
	}

	return nil
}
