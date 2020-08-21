// +build integration
// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/paingha/winkel/api/config"
	"github.com/paingha/winkel/api/models"
	"github.com/jinzhu/gorm"

	"github.com/paingha/winkel/api/routes"
	"github.com/stretchr/testify/assert"
)

func init() {
	var err error
	config.DB, err = gorm.Open("postgres", config.GetConnectionContext())
	if err != nil {
		fmt.Printf("connection error: %s", err)
	}
	config.DB.LogMode(true)
	config.DB.AutoMigrate(&models.User{})
}
func TestGetUserHandler(t *testing.T) {
	testRouter := routes.SetupRouter()

	req, err := http.NewRequest("GET", "/v1/user/9", nil)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	assert.Equal(t, 200, resp.Code)
}

/*
func TestGetUsersHandler(t *testing.T) {
	var err error
	t.Log("test")
	config.DB, err = gorm.Open("postgres", config.GetConnectionContext())
	if err != nil {
		fmt.Printf("connection error: %s", err)
	}
	defer config.DB.Close()
	// run the migrations: User struct
	config.DB.LogMode(true)
	config.DB.AutoMigrate(&models.User{})
	testRouter := routes.SetupRouter()

	req, err := http.NewRequest("GET", "/v1/user?offset=0&limit=12", nil)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	t.Log(resp)
	assert.Equal(t, 200, resp.Code)
}
*/

func TestLoginUserRightHandler(t *testing.T) {
	t.Log("test")
	testRouter := routes.SetupRouter()
	body, _ := json.Marshal(map[string]string{
		"email":    "apaingha@gmail.com",
		"password": "Computer2",
	})
	Body := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", "/v1/user/login", Body)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	t.Log(resp)
	assert.Equal(t, 200, resp.Code)
}

func TestLoginUserWrongHandler(t *testing.T) {
	testRouter := routes.SetupRouter()
	body, _ := json.Marshal(map[string]string{
		"email":    "apaingha@gmail.com",
		"password": "Computer2!",
	})
	Body := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", "/v1/user/login", Body)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	t.Log(resp)
	assert.Equal(t, 400, resp.Code)
}

func TestRegisterUserSuccessHandler(t *testing.T) {
	testRouter := routes.SetupRouter()
	body, _ := json.Marshal(map[string]string{
		"email":     "test11@test.com",
		"password":  "Computer2!",
		"firstName": "Testing",
		"lastName":  "Test",
	})
	Body := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", "/v1/user/register", Body)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	t.Log(resp)
	assert.Equal(t, 200, resp.Code)
}

func TestRegisterUserFailureHandler(t *testing.T) {
	testRouter := routes.SetupRouter()
	body, _ := json.Marshal(map[string]string{
		"email":     "apaingha.com",
		"password":  "Computer2!",
		"firstName": "Testing",
		"lastName":  "Test",
	})
	Body := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", "/v1/user/register", Body)
	if err != nil {
		fmt.Println(err)
	}

	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	t.Log(resp)
	assert.Equal(t, 409, resp.Code)
	defer config.DB.Close()
}
