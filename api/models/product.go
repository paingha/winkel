// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import (
	"time"

	"github.com/paingha/winkel/api/config"
)

//Products model struct
type Products struct {
	ID            int        `json:"id,omitempty" sql:"primary_key"`
	Name          string     `json:"name"`
	UUID          string     `gorm:"unique" json:"uuid"`
	Description   string     `json:"description"`
	FeaturedImage string     `json:"featuredImage"`
	Price         float64    `json:"price"`
	Stock         int        `json:"stock"`
	CategoryID    int        `gorm:"not null" json:"categoryId"`
	SubCategoryID int        `gorm:"not null" json:"subcategoryId"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `json:"deletedAt"`
}

//TableName - products db table name override
func (Products) TableName() string {
	return "products"
}

//GetAllProducts - fetch all products at once
func GetAllProducts(product *[]Products, offset, limit int) (int, error) {
	var count = 0
	if err := config.DB.Model(&Products{}).Count(&count).Order("created_at desc").Offset(offset).Limit(limit).Find(product).Error; err != nil {
		return count, err
	}
	return count, nil
}

//GetProductByCategoryAndSubCategory - fetch all products according to category id and sub category id
func GetProductByCategoryAndSubCategory(product *[]Products, categoryID, subCategoryID, offset, limit int) (int, error) {
	var count = 0
	if err := config.DB.Model(&Products{}).Count(&count).Order("created_at desc").Offset(offset).Limit(limit).Where(&Products{CategoryID: categoryID, SubCategoryID: subCategoryID}).Find(product).Error; err != nil {
		return count, err
	}
	return count, nil
}

//CreateProduct - create a product
func CreateProduct(product *Products) (bool, error) {
	if err := config.DB.Create(product).Error; err != nil {
		return false, err
	}
	return false, nil
}

//GetProduct - fetch one product
func GetProduct(product *Products, id int) error {
	if err := config.DB.Where("id = ?", id).First(product).Error; err != nil {
		return err
	}
	return nil
}

//UpdateProduct - update a product
func UpdateProduct(product *Products, id int) error {
	if err := config.DB.Model(&product).Where("id = ?", id).Updates(product).Error; err != nil {
		return err
	}
	return nil
}

//DeleteProduct - delete a product
func DeleteProduct(id int) error {
	if err := config.DB.Where("id = ?", id).Unscoped().Delete(Products{}).Error; err != nil {
		return err
	}
	return nil
}
