// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/paingha/winkel/api/controllers"
	"github.com/paingha/winkel/api/middlewares"
)

//CategoryRouter - category subroutes
func CategoryRouter(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/v1/category")
	{
		v1.POST("/create", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.CategoryControllers["createCategory"])
		v1.GET("/:id", middlewares.AuthenticationMiddleware(), controllers.CategoryControllers["getCategory"])
		v1.PUT("/:id", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.CategoryControllers["updateCategory"])
		v1.DELETE("/:id", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.CategoryControllers["deleteCategory"])
		v1.GET("/", middlewares.AuthenticationMiddleware(), controllers.CategoryControllers["getCategories"])

	}
	return r
}
