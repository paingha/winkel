// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils_test

import (
	"testing"

	"github.com/paingha/winkel/api/utils"
	"github.com/stretchr/testify/assert"
)

func TestConvertstringtoint(T *testing.T) {
	resp, err := utils.ConvertStringToInt("100")
	assert.Nil(T, err)
	assert.Equal(T, resp, 100, "The returned integer should be 100")
}
