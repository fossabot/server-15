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

import "github.com/gofiber/fiber/v2/utils"

func nilUint8(i uint8) *uint8 {
	if i == 0 {
		return nil
	}

	return &i
}

func nilUint16(i uint16) *uint16 {
	if i == 0 {
		return nil
	}

	return &i
}

func nilString(s string) *string {
	if s == "" {
		return nil
	}

	s = utils.CopyString(s)

	return &s
}
