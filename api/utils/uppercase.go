// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"strings"
)

//UppercaseName convert user first name and last name to have first character capitalized
func UppercaseName(name string) string {
	var emptyString []string
	newString := append(emptyString, strings.ToUpper(name[:1]), name[1:])
	//Join the array of string to make a single string
	concatedString := strings.Join(newString, "")
	return concatedString
}
