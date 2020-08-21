// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

//GenerateRandomString generates random string of set length
import (
	"math/rand"
	"time"
)

const numbers = "123456789"

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

const numcharset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

//StringWithCharset - takes in a group of characters and used the seeded random function to pick randomly
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

//GenerateRandomString - generates a random string of provided length
func GenerateRandomString(n int) string {
	return StringWithCharset(n, charset)
}

//GenerateRandomInt - generates a random int string of provided length
func GenerateRandomInt(n int) string {
	return StringWithCharset(n, numbers)
}

func NumWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = numcharset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GenerateRandomNum(n int) string {
	return NumWithCharset(n, charset)
}
