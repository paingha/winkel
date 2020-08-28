// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
)

//SetupRouter - setup routes for api
func SetupRouter() *gin.Engine {
	r := gin.Default()
	{
		RootRouter(r)
		UserRouter(r)
		CategoryRouter(r)
		SubCategoryRouter(r)
		ProductRouter(r)
	}
	return r
}
