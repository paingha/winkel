// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package middlewares

import (
	"net/http"

	"github.com/paingha/winkel/api/security"
	"github.com/gin-gonic/gin"
)

//AdminMiddleware - Auth guard middleware
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		s, good := c.Get("session")
		session, very := s.(*security.Claims)
		if !session.IsAdmin || !good || !very {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message":    "Unauthorized",
				"statusCode": 409,
			})
		}
		c.Next()
	}
}
