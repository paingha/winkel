// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import (
	"time"

	"github.com/paingha/winkel/api/config"
)

//Subcategories - subcategory data struct
type Subcategories struct {
	ID            int        `json:"id,omitempty" sql:"primary_key"`
	Name          string     `gorm:"not null" json:"name"`
	Description   string     `json:"description"`
	FeaturedImage string     `json:"featuredImage"`
	Icon          string     `json:"icon"`
	CategoryID    int        `gorm:"not null" json:"categoryId"`
	CreatedAt     time.Time  `json:"createdAt"`
	UpdatedAt     time.Time  `json:"updatedAt"`
	DeletedAt     *time.Time `json:"deletedAt"`
}

//TableName - subcategories db table name override
func (Subcategories) TableName() string {
	return "subcategories"
}

//GetAllSubcategories - fetch all subcategories at once
func GetAllSubcategories(subcategory *[]Subcategories, offset int, limit int) (counts int, err error) {
	var count = 0
	if err = config.DB.Model(&Subcategories{}).Count(&count).Order("created_at desc").Offset(offset).Limit(limit).Find(subcategory).Error; err != nil {
		return count, err
	}
	return count, nil
}

//GetCategorySubcategories - fetch conversation subcategories
func GetCategorySubcategories(subcategory *[]Subcategories, id int, offset int, limit int) (counts int, err error) {
	var count = 0
	if err = config.DB.Model(&Subcategories{}).Where("category_id = ?", id).Count(&count).Order("created_at desc").Offset(offset).Limit(limit).Find(subcategory).Error; err != nil {
		return count, err
	}
	return count, nil
}

//CreateSubcategory - create a subcategory
func CreateSubcategory(subcategory *Subcategories) (created bool, err error) {
	if err := config.DB.Create(subcategory).Error; err != nil {
		return false, err
	}
	return false, nil
}

//GetSubcategory - fetch one subcategory
func GetSubcategory(subcategory *Subcategories, id int) (err error) {
	if err := config.DB.Where("id = ?", id).First(subcategory).Error; err != nil {
		return err
	}
	return nil
}

//UpdateSubcategory - update a subcategory
func UpdateSubcategory(subcategory *Subcategories, id int) (err error) {
	config.DB.Model(&subcategory).Where("id = ?", id).Updates(subcategory)
	return nil
}

//DeleteSubcategory - delete a subcategory
func DeleteSubcategory(id int) (err error) {
	if err := config.DB.Where("id = ?", id).Unscoped().Delete(Subcategories{}).Error; err != nil {
		return err
	}
	return nil
}
