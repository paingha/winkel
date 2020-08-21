// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/paingha/winkel/api/models"
	"github.com/paingha/winkel/api/plugins"
	"github.com/paingha/winkel/api/utils"
)

//CategoryControllers - map of all the categories controllers
var CategoryControllers = map[string]func(*gin.Context){
	"getCategories":  GetCategories,
	"createCategory": CreateCategory,
	"getCategory":    GetCategory,
	"updateCategory": UpdateCategory,
	"deleteCategory": DeleteCategory,
}

//GetCategories - List all Categories
// @Summary List all Categories
// @Tags Category
// @Produce json
// @Success 200 {object} models.Categories
// @Router /category [get]
// @Security ApiKeyAuth
func GetCategories(c *gin.Context) {
	offsetString := c.Query("offset")
	limitString := c.Query("limit")
	offset, err := utils.ConvertStringToInt(offsetString)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Offset conv error",
			"statusCode": 400,
		})
		plugins.LogError("API", "Offset conv error", err)
	}
	limit, errs := utils.ConvertStringToInt(limitString)
	if errs != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Limit conv error",
			"statusCode": 400,
		})
		plugins.LogError("API", "Limit conv error", errs)
	}
	var category []models.Categories
	count, err := models.GetAllCategories(&category, offset, limit)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error getting all categories", err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"count":      count,
			"data":       category,
			"statusCode": 200,
		})
	}
}

//CreateCategory - Create a Category
// @Summary Creates a new Category
// @Description Creates a new Category
// @Tags Category
// @Accept  json
// @Produce json
// @Param category body models.Categories true "Create Category"
// @Success 200 {object} models.Categories
// @Router /category/create [post]
// @Security ApiKeyAuth
func CreateCategory(c *gin.Context) {
	var category models.Categories
	c.BindJSON(&category)
	stats, err := models.CreateCategory(&category)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error creating category", err)
	} else {
		if stats != true {
			c.JSON(http.StatusConflict, gin.H{
				"message":    "Category already exists",
				"statusCode": 409,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message":    "Category created successfully",
				"statusCode": 200,
			})
		}
	}
}

//GetCategory - Get a particular Category with id
// @Summary Retrieves category based on given ID
// @Tags Category
// @Produce json
// @Param id path integer true "Category ID"
// @Success 200 {object} models.Categories
// @Router /category/{id} [get]
// @Security ApiKeyAuth
func GetCategory(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, idErr := utils.ConvertStringToInt(idString)
	if idErr != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Conv error",
			"statusCode": 400,
		})
		plugins.LogError("API", "error id conv to int", idErr)
	}
	var category models.Categories
	err := models.GetCategory(&category, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		plugins.LogError("API", "error getting category", err)
	} else {
		c.JSON(http.StatusOK, category)
	}

}

//UpdateCategory - Update an existing Category
// @Summary Updates category based on given ID
// @Tags Category
// @Produce json
// @Param id path integer true "Category ID"
// @Success 200 {object} models.Categories
// @Router /category/{id} [patch]
// @Security ApiKeyAuth
func UpdateCategory(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, idErr := utils.ConvertStringToInt(idString)
	if idErr != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Conv error",
			"statusCode": 400,
		})
		plugins.LogError("API", "error id conv to int", idErr)
	}
	var category models.Categories
	err := models.GetCategory(&category, id)
	if err != nil {
		c.JSON(http.StatusNotFound, category)
		plugins.LogError("API", "error getting category", err)
	}
	c.BindJSON(&category)
	err = models.UpdateCategory(&category, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error updating category", err)
	} else {
		c.JSON(http.StatusOK, category)
	}

}

//DeleteCategory - Deletes Category
// @Summary Deletes a category based on given ID
// @Tags Category
// @Produce json
// @Param id path integer true "Category ID"
// @Success 200 {object} models.Categories
// @Router /category/{id} [delete]
// @Security ApiKeyAuth
func DeleteCategory(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, err := utils.ConvertStringToInt(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":    "Id conv error",
			"statusCode": 500,
		})
		plugins.LogError("API", "error id conv to int", err)
	}
	errs := models.DeleteCategory(id)
	if errs != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error deleting category", errs)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":    "Category Deleted successfully",
			"statusCode": 200,
		})
	}
}
