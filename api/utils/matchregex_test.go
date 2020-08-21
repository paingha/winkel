// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils_test

import (
	"testing"

	"github.com/paingha/winkel/api/utils"
	"github.com/stretchr/testify/assert"
)

func TestMatchregex(T *testing.T) {
	resp, err := utils.MatchRegex("joe", "joey alagoa")
	assert.Nil(T, err, "There should be no error returned")
	assert.Equal(T, resp, true, "The name should be equal")
}
func TestFalseMatchregex(T *testing.T) {
	resp, err := utils.MatchRegex("joe", "oej")
	assert.Nil(T, err, "There should be no error returned")
	assert.Equal(T, resp, false, "The name should be equal")
}
