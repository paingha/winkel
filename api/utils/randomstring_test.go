// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils_test

import (
	"testing"

	"github.com/paingha/winkel/api/utils"
	"github.com/stretchr/testify/assert"
)

func TestRandomstring(T *testing.T) {
	resp := utils.GenerateRandomString(10)
	assert.NotNil(T, resp)
	assert.Equal(T, len(resp), 10, "The returned string length should be 10")
}

func TestRandomInt(T *testing.T) {
	resp := utils.GenerateRandomInt(5)
	assert.NotNil(T, resp)
	assert.Equal(T, len(resp), 5, "The returned int string length should be 5")
}
