// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"strconv"
)

//ConvertStringToInt - converts a string to an integer
func ConvertStringToInt(character string) (data int, err error) {
	i, errs := strconv.Atoi(character)
	return i, errs
}
