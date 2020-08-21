// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package security

import (
	"os"

	"github.com/paingha/winkel/api/plugins"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

//Claims struct for Jwt
type Claims struct {
	ID      string `json:"id,omitmpty"`
	UserID  int    `json:"user_id,omitmpty"`
	IsAdmin bool   `json:"isAdmin,omitmpty"`
	jwt.StandardClaims
}

//CreateJWT - creates json web token for user login
func CreateJWT(userID int, admin bool) (string, error) {
	u2, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	claims := Claims{
		ID:      u2.String(),
		UserID:  userID,
		IsAdmin: admin,
	}
	jwtSecretByte := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretByte)
}

//VerifyJWT takes in token as a string and returns a boolean.
func VerifyJWT(jwtToken string) (*Claims, bool) {
	// Initialize a new instance of `Claims`
	//fmt.Println(jwtToken)
	claims := &Claims{}

	// Parse the JWT string and store the result in `claims`.
	// Note that we are passing the key in this method as well. This method will return an error
	// if the token is invalid (if it has expired according to the expiry time we set on sign in),
	// or if the signature does not match
	tkn, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			plugins.LogError("API", "error jwt signature is invalid", err)
			return nil, false
		}
		plugins.LogError("API", "error bad jwt request", err)
		return nil, false
	}
	if !tkn.Valid {
		plugins.LogError("API", "inavlid jwt token", err)
		return nil, false
	}
	// Finally, return the welcome message to the user, along with their
	// username given in the token
	return claims, true
}
