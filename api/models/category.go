// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import (
	"time"

	"github.com/paingha/winkel/api/config"
)

//Categories model struct
type Categories struct {
	ID        int        `json:"id,omitempty" sql:"primary_key"`
	Name      string     `json:"name"`
	UUID      string     `gorm:"unique" json:"uuid"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

//TableName - categories db table name override
func (Categories) TableName() string {
	return "categories"
}

//GetAllCategories - fetch all categories at once
func GetAllCategories(category *[]Categories, offset int, limit int) (counts int, err error) {
	var count = 0
	if err = config.DB.Model(&Categories{}).Count(&count).Order("created_at desc").Offset(offset).Limit(limit).Find(category).Error; err != nil {
		return count, err
	}
	return count, nil
}

//CreateCategory - create a category
func CreateCategory(category *Categories) (created bool, err error) {
	if err := config.DB.Create(category).Error; err != nil {
		return false, err
	}
	return false, nil
}

//GetCategory - fetch one category
func GetCategory(category *Categories, id int) (err error) {
	if err := config.DB.Where("id = ?", id).First(category).Error; err != nil {
		return err
	}
	return nil
}

//UpdateCategory - update a category
func UpdateCategory(category *Categories, id int) (err error) {
	config.DB.Model(&category).Where("id = ?", id).Updates(category)
	return nil
}

//DeleteCategory - delete a category
func DeleteCategory(id int) (err error) {
	if err := config.DB.Where("id = ?", id).Unscoped().Delete(Categories{}).Error; err != nil {
		return err
	}
	return nil
}
