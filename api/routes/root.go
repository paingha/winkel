// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//RootRouter - root subroutes
func RootRouter(r *gin.Engine) *gin.Engine {
	r.Static("/public", "./public")
	root := r.Group("/")
	v1 := r.Group("/v1")
	r.LoadHTMLGlob("templates/*.html")
	{
		root.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
		v1.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
	}
	return r
}
