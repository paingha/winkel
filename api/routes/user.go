// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/paingha/winkel/api/controllers"
	"github.com/paingha/winkel/api/middlewares"
	"github.com/gin-gonic/gin"
)

//UserRouter - user subroutes
func UserRouter(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/v1/user")
	{
		v1.POST("/verify-phone/:id", middlewares.AuthenticationMiddleware(), controllers.UserControllers["verifyPhoneUser"])
		v1.POST("/verify-phone-code/:id", middlewares.AuthenticationMiddleware(), controllers.UserControllers["verifyPhoneCodeUser"])
		v1.GET("/:id", middlewares.AuthenticationMiddleware(), controllers.UserControllers["getUser"])
		v1.PUT("/:id", middlewares.AuthenticationMiddleware(), controllers.UserControllers["updateUsers"])
		v1.DELETE("/:id", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.UserControllers["deleteUser"])
		v1.GET("/", middlewares.AuthenticationMiddleware(), middlewares.AdminMiddleware(), controllers.UserControllers["getUsers"])
		v1.POST("/register", controllers.UserControllers["createUser"])
		v1.POST("/login", controllers.UserControllers["loginUser"])
		v1.GET("/:id/verify-email", controllers.UserControllers["verifyEmailUser"])
	}
	return r
}
