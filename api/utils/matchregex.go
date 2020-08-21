// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"regexp"
)

//MatchRegex function to test if a string matches another string using regex. Returns a bool and an error if present
func MatchRegex(matchMe string, shouldMatch string) (bool, error) {
	//Concantenate the string to the perl version of matching numbers and strings
	placeholder := matchMe
	compiledPointer, err := regexp.Compile(placeholder)
	if err != nil {
		return false, err
	}
	result := compiledPointer.Match([]byte(shouldMatch))
	//if it gets here it means err will be <nil>
	return result, err
}
