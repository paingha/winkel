// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/paingha/winkel/api/controllers"
	"github.com/paingha/winkel/api/middlewares"
)

//ProductRouter - product subroutes
func ProductRouter(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/v1/product")
	{
		v1.POST("/create", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.ProductControllers["createProduct"])
		v1.GET("/:id", middlewares.AuthenticationMiddleware(), controllers.ProductControllers["getProduct"])
		v1.GET("/:id/category/:id/subcategory", middlewares.AuthenticationMiddleware(), controllers.ProductControllers["getProductByCategoryAndSubCategory"])
		v1.PUT("/:id", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.ProductControllers["updateProduct"])
		v1.DELETE("/:id", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.ProductControllers["deleteProduct"])
		v1.GET("/", middlewares.AuthenticationMiddleware(), controllers.ProductControllers["getProducts"])

	}
	return r
}
