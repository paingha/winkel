// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/paingha/winkel/api/controllers"
	"github.com/paingha/winkel/api/middlewares"
)

//SubCategoryRouter - subcategory subroutes
func SubCategoryRouter(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/v1/sub-category")
	{
		v1.POST("/create", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.SubcategoryControllers["createSubcategory"])
		v1.GET("/:id", middlewares.AuthenticationMiddleware(), controllers.SubcategoryControllers["getSubcategory"])
		v1.GET("/:id/category", middlewares.AuthenticationMiddleware(), controllers.SubcategoryControllers["getCategorySubcategories"])
		v1.PUT("/:id", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.SubcategoryControllers["updateSubcategory"])
		v1.DELETE("/:id", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.SubcategoryControllers["deleteSubcategory"])
		v1.GET("/", middlewares.AuthenticationMiddleware(), controllers.SubcategoryControllers["getSubcategories"])

	}
	return r
}
