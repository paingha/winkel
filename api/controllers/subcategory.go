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

//SubcategoryControllers - map of all the subcategories controllers
var SubcategoryControllers = map[string]func(*gin.Context){
	"getSubcategories":         GetSubcategories,
	"getCategorySubcategories": GetCategorySubcategories,
	"createSubcategory":        CreateSubcategory,
	"getSubcategory":           GetSubcategory,
	"updateSubcategory":        UpdateSubcategory,
	"deleteSubcategory":        DeleteSubcategory,
}

//GetSubcategories - List all Subcategories
// @Summary List all Subcategories
// @Tags Subcategory
// @Produce json
// @Success 200 {object} models.Subcategories
// @Router /subcategory [get]
// @Security ApiKeyAuth
func GetSubcategories(c *gin.Context) {
	offsetString := c.Query("offset")
	limitString := c.Query("limit")
	offset, err := utils.ConvertStringToInt(offsetString)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"subcategory": "Offset conv error",
			"statusCode":  400,
		})
		plugins.LogError("API", "error offset conv to int", err)
	}
	limit, errs := utils.ConvertStringToInt(limitString)
	if errs != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"subcategory": "Limit conv error",
			"statusCode":  400,
		})
		plugins.LogError("API", "error limit conv to int", errs)
	}
	var subcategory []models.Subcategories
	count, err := models.GetAllSubcategories(&subcategory, offset, limit)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error getting subcategories", err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"count":      count,
			"data":       subcategory,
			"statusCode": 200,
		})
	}
}

//GetCategorySubcategories - List all Conversation Subcategories
// @Summary List all my Subcategories
// @Tags Subcategory
// @Produce json
// @Success 200 {object} models.Subcategories
// @Router /subcategory/:id/my [get]
// @Security ApiKeyAuth
func GetCategorySubcategories(c *gin.Context) {
	offsetString := c.Query("offset")
	limitString := c.Query("limit")
	idString := c.Params.ByName("id")
	id, idErr := utils.ConvertStringToInt(idString)
	offset, offsetErr := utils.ConvertStringToInt(offsetString)
	limit, errs := utils.ConvertStringToInt(limitString)
	if offsetErr != nil || idErr != nil || errs != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"subcategory": "Conv error",
			"statusCode":  400,
		})
		plugins.LogError("API", "error offset conv to int", offsetErr)
		plugins.LogError("API", "error id conv to int", idErr)
		plugins.LogError("API", "error limit conv to int", errs)
	}

	var subcategory []models.Subcategories
	count, err := models.GetCategorySubcategories(&subcategory, id, offset, limit)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error getting category sub categories", err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"count":      count,
			"data":       subcategory,
			"statusCode": 200,
		})
	}

}

//CreateSubcategory - Create a Subcategory
// @Summary Creates a new Subcategory
// @Description Creates a new Subcategory
// @Tags Subcategory
// @Accept  json
// @Produce json
// @Param subcategory body models.Subcategories true "Create Subcategory"
// @Success 200 {object} models.Subcategories
// @Router /subcategory/create [post]
// @Security ApiKeyAuth
func CreateSubcategory(c *gin.Context) {
	var subcategory models.Subcategories
	c.BindJSON(&subcategory)
	stats, err := models.CreateSubcategory(&subcategory)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error creating subcategory", err)
	} else {
		if stats != true {
			c.JSON(http.StatusConflict, gin.H{
				"subcategory": "Subcategory already exists",
				"statusCode":  409,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"subcategory": "Subcategory created successfully",
				"statusCode":  200,
			})
		}
	}
}

//GetSubcategory - Get a particular Subcategory with id
// @Summary Retrieves subcategory based on given ID
// @Tags Subcategory
// @Produce json
// @Param id path integer true "Subcategory ID"
// @Success 200 {object} models.Subcategories
// @Router /subcategory/{id} [get]
// @Security ApiKeyAuth
func GetSubcategory(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, idErr := utils.ConvertStringToInt(idString)
	if idErr != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"subcategory": "Conv error",
			"statusCode":  400,
		})
		plugins.LogError("API", "error id conv to int", idErr)
	}
	var subcategory models.Subcategories
	err := models.GetSubcategory(&subcategory, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		plugins.LogError("API", "error getting subcategory", err)
	} else {
		c.JSON(http.StatusOK, subcategory)
	}

}

//UpdateSubcategory - Update an existing Subcategory
// @Summary Updates subcategory based on given ID
// @Tags Subcategory
// @Produce json
// @Param id path integer true "Subcategory ID"
// @Success 200 {object} models.Subcategories
// @Router /subcategory/{id} [patch]
// @Security ApiKeyAuth
func UpdateSubcategory(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, idErr := utils.ConvertStringToInt(idString)
	if idErr != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"subcategory": "Conv error",
			"statusCode":  400,
		})
		plugins.LogError("API", "error id conv to int", idErr)
	}
	var subcategory models.Subcategories
	err := models.GetSubcategory(&subcategory, id)
	if err != nil {
		c.JSON(http.StatusNotFound, subcategory)
		plugins.LogError("API", "error getting subcategory", err)
	}
	c.BindJSON(&subcategory)
	err = models.UpdateSubcategory(&subcategory, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error updating subcategory", err)
	} else {
		c.JSON(http.StatusOK, subcategory)
	}

}

//DeleteSubcategory - Deletes Subcategory
// @Summary Deletes a subcategory based on given ID
// @Tags Subcategory
// @Produce json
// @Param id path integer true "Subcategory ID"
// @Success 200 {object} models.Subcategories
// @Router /subcategory/{id} [delete]
// @Security ApiKeyAuth
func DeleteSubcategory(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, err := utils.ConvertStringToInt(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"subcategory": "Id conv error",
			"statusCode":  500,
		})
		plugins.LogError("API", "error id conv to int", err)
	}
	errs := models.DeleteSubcategory(id)
	if errs != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error deleting subcategory. subcategory not found", errs)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"subcategory": "Subcategory Deleted successfully",
			"statusCode":  200,
		})
	}
}
