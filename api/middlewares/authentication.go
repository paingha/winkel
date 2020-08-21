// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middlewares

import (
	"net/http"
	"strings"

	"github.com/paingha/winkel/api/security"
	"github.com/gin-gonic/gin"
)

//AuthenticationMiddleware - Auth guard middleware
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		token := strings.TrimPrefix(header, "Bearer ")
		claims, ok := security.VerifyJWT(token)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
				"message":    "Unauthenticated",
				"statusCode": 401,
			})
		}
		c.Set("session", claims)
		c.Next()
	}
}

//GetSession - returns claims from header session
func GetSession(c *gin.Context) (*security.Claims, bool) {
	s, ok := c.Get("session")
	if !ok {
		return nil, false
	}
	session, ok := s.(*security.Claims)
	if !ok {
		return nil, false
	}
	return session, true
}
