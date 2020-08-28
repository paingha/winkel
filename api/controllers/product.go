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

//ProductControllers - map of all the products controllers
var ProductControllers = map[string]func(*gin.Context){
	"getProducts":                        GetProducts,
	"getProductByCategoryAndSubCategory": GetProductByCategoryAndSubCategory,
	"createProduct":                      CreateProduct,
	"getProduct":                         GetProduct,
	"updateProduct":                      UpdateProduct,
	"deleteProduct":                      DeleteProduct,
}

//GetProducts - List all Products
// @Summary List all Products
// @Tags Product
// @Produce json
// @Success 200 {object} models.Products
// @Router /product [get]
// @Security ApiKeyAuth
func GetProducts(c *gin.Context) {
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
	var product []models.Products
	count, err := models.GetAllProducts(&product, offset, limit)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error getting all products", err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"count":      count,
			"data":       product,
			"statusCode": 200,
		})
	}
}

//CreateProduct - Create a Product
// @Summary Creates a new Product
// @Description Creates a new Product
// @Tags Product
// @Accept  json
// @Produce json
// @Param product body models.Products true "Create Product"
// @Success 200 {object} models.Products
// @Router /product/create [post]
// @Security ApiKeyAuth
func CreateProduct(c *gin.Context) {
	var product models.Products
	c.BindJSON(&product)
	stats, err := models.CreateProduct(&product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error creating product", err)
	} else {
		if stats != true {
			c.JSON(http.StatusConflict, gin.H{
				"message":    "Product already exists",
				"statusCode": 409,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message":    "Product created successfully",
				"statusCode": 200,
			})
		}
	}
}

//GetProduct - Get a particular Product with id
// @Summary Retrieves product based on given ID
// @Tags Product
// @Produce json
// @Param id path integer true "Product ID"
// @Success 200 {object} models.Products
// @Router /product/{id} [get]
// @Security ApiKeyAuth
func GetProduct(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, idErr := utils.ConvertStringToInt(idString)
	if idErr != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Conv error",
			"statusCode": 400,
		})
		plugins.LogError("API", "error id conv to int", idErr)
	}
	var product models.Products
	err := models.GetProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		plugins.LogError("API", "error getting product", err)
	} else {
		c.JSON(http.StatusOK, product)
	}

}

//GetProductByCategoryAndSubCategory - List all Products by Category and Sub Category
// @Summary List all Products by Category and Sub Category
// @Tags Product
// @Produce json
// @Success 200 {object} models.Products
// @Router /product/by-catgeory [get]
// @Security ApiKeyAuth
func GetProductByCategoryAndSubCategory(c *gin.Context) {
	offsetString := c.Query("offset")
	limitString := c.Query("limit")
	categoryIDString := c.Params.ByName("id")
	subcategoryIDString := c.Params.ByName("id")
	offset, err := utils.ConvertStringToInt(offsetString)
	catgeoryID, catgeoryIDErr := utils.ConvertStringToInt(categoryIDString)
	subCategoryID, subCategoryIDErr := utils.ConvertStringToInt(subcategoryIDString)
	if catgeoryIDErr != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Conv error",
			"statusCode": 400,
		})
		plugins.LogError("API", "error catgeory ID conv to int", catgeoryIDErr)
	}
	if subCategoryIDErr != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Conv error",
			"statusCode": 400,
		})
		plugins.LogError("API", "error Sub Catgeory ID conv to int", subCategoryIDErr)
	}
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
	var product []models.Products
	count, err := models.GetProductByCategoryAndSubCategory(&product, catgeoryID, subCategoryID, offset, limit)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error getting all products by category and sub category", err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"count":      count,
			"data":       product,
			"statusCode": 200,
		})
	}
}

//UpdateProduct - Update an existing Product
// @Summary Updates product based on given ID
// @Tags Product
// @Produce json
// @Param id path integer true "Product ID"
// @Success 200 {object} models.Products
// @Router /product/{id} [patch]
// @Security ApiKeyAuth
func UpdateProduct(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, idErr := utils.ConvertStringToInt(idString)
	if idErr != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Conv error",
			"statusCode": 400,
		})
		plugins.LogError("API", "error id conv to int", idErr)
	}
	var product models.Products
	err := models.GetProduct(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
		plugins.LogError("API", "error getting product", err)
	}
	c.BindJSON(&product)
	err = models.UpdateProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error updating product", err)
	} else {
		c.JSON(http.StatusOK, product)
	}

}

//DeleteProduct - Deletes Product
// @Summary Deletes a product based on given ID
// @Tags Product
// @Produce json
// @Param id path integer true "Product ID"
// @Success 200 {object} models.Products
// @Router /product/{id} [delete]
// @Security ApiKeyAuth
func DeleteProduct(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, err := utils.ConvertStringToInt(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":    "Id conv error",
			"statusCode": 500,
		})
		plugins.LogError("API", "error id conv to int", err)
	}
	errs := models.DeleteProduct(id)
	if errs != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "error deleting product", errs)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":    "Product Deleted successfully",
			"statusCode": 200,
		})
	}
}
