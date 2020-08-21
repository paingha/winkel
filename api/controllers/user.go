// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package controllers

import (
	"encoding/base64"
	"net/http"

	"github.com/paingha/winkel/api/middlewares"
	"github.com/paingha/winkel/api/models"
	"github.com/paingha/winkel/api/plugins"
	"github.com/paingha/winkel/api/security"
	"github.com/paingha/winkel/api/utils"
	"github.com/gin-gonic/gin"
)

//UserControllers - map of all the users controllers
var UserControllers = map[string]func(*gin.Context){
	"getUsers":            GetUsers,
	"createUser":          CreateUser,
	"loginUser":           LoginUser,
	"getUser":             GetUser,
	"verifyEmailUser":     VerifyEmailUser,
	"verifyPhoneUser":     VerifyPhoneUser,
	"verifyPhoneCodeUser": VerifyPhoneCodeUser,
	"forgotUser":          ForgotUser,
	"updateUser":          UpdateUser,
	"deleteUser":          DeleteUser,
}

//GetUsers - List all Users
// @Summary List all registered Users
// @Tags User Auth
// @Produce json
// @Success 200 {object} models.User
// @Router /user [get]
// @Security ApiKeyAuth
func GetUsers(c *gin.Context) {
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
	var user []models.User
	count, err := models.GetAllUsers(&user, offset, limit)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "Error getting all users", err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"count":      count,
			"data":       user,
			"statusCode": 200,
		})
	}
}

//CreateUser - Create a User
// @Summary Registers a new User
// @Description Creates a new User account
// @Tags User Auth
// @Accept  json
// @Produce json
// @Param user body models.User true "Create User Account"
// @Success 200 {object} models.User
// @Router /user/register [post]
func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	//Hash the user's password for security
	user.Password = security.HashSaltPassword([]byte(user.Password))
	//capitalize the first letter of the User's first and last name
	user.FirstName = utils.UppercaseName(user.FirstName)
	user.LastName = utils.UppercaseName(user.LastName)
	//end here
	//Generate a random string for the verify code
	user.VerifyCode = utils.GenerateRandomString(30)
	//end here
	stats, err := models.CreateUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "Error unable to create user", err)
	} else {
		if stats != true {
			c.JSON(http.StatusConflict, gin.H{
				"message":    "Account already exists",
				"statusCode": 409,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message":    "Account created successfully",
				"statusCode": 200,
			})
		}
	}
}

//LoginUser - Login a User
// @Summary Logins a User
// @Description Login a user by sending jwt
// @Tags User Auth
// @Accept  json
// @Produce json
// @Success 200 {object} models.User
// @Router /user/login [post]
func LoginUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	users, token, err := models.LoginUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		plugins.LogError("API", "Login error", err)
	} else {
		if token == "" && err == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message":    "Invalid Credentials",
				"statusCode": 400,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"token":      token,
				"data":       users,
				"statusCode": 200,
			})
		}
	}
}

//ForgotUser - Forgot User Password
// @Summary Forgot User Password Endpoint
// @Description Forgot User Password Endpoint
// @Tags User Auth
// @Accept  json
// @Produce json
// @Success 200 {object} models.User
// @Router /user/forgot [post]
func ForgotUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	state, err := models.ForgotUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		plugins.LogError("API", "Forgot User error", err)
	} else {
		if !state && err == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message":    "An error occured",
				"statusCode": 400,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"data":       "Password Reset Email sent successfully",
				"statusCode": 200,
			})
		}
	}
}

//GetUser - Get a particular User with id
// @Summary Retrieves user based on given ID
// @Tags User Auth
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /user/{id} [get]
// @Security ApiKeyAuth
func GetUser(c *gin.Context) {
	idString := c.Params.ByName("id")
	var user models.User
	res, _ := middlewares.GetSession(c)
	id, err := utils.ConvertStringToInt(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":    "Id conv error",
			"statusCode": 500,
		})
		plugins.LogError("API", "ID conv error", err)
	}
	if res.IsAdmin || id == res.UserID {
		err := models.GetUser(&user, id)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			plugins.LogError("API", "Get User error", err)
		} else {
			c.JSON(http.StatusOK, user)
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Needs Elevation",
			"statusCode": 400,
		})
	}
}

//VerifyEmailUser - Get a particular User with id
// @Summary Verifies a user's email
// @Tags User Auth
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /user/verify-email [post]
//Router /user/verify-email?token={id} [post]
func VerifyEmailUser(c *gin.Context) {
	token := c.Query("token")
	var user models.User
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Token error",
			"statusCode": 400,
		})
		return
	}
	if tokenDecoded, err := base64.StdEncoding.DecodeString(token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Token error",
			"statusCode": 400,
		})
		plugins.LogError("API", "Email verify token decode error", err)
	} else {
		token = string(tokenDecoded)
	}
	err := models.VerifyEmailUser(&user, token)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		plugins.LogError("API", "Error verifying User Email", err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":    "Email verified successfully",
			"statusCode": 200,
		})
	}

}

//VerifyPhoneUser - Sends Phone Verification Code to User
// @Summary Send phone code to verify User's phone number
// @Description Verifies user by sending verification code to user through whatsapp or sms
// @Tags User Auth
// @Accept  json
// @Produce json
// @Success 200 {object} models.User
// @Router /user/verify-phone/:id [post]
// @Security ApiKeyAuth
func VerifyPhoneUser(c *gin.Context) {
	var user models.User
	medium := c.Query("medium")
	idString := c.Params.ByName("id")
	id, err := utils.ConvertStringToInt(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":    "Id conv error",
			"statusCode": 500,
		})
		plugins.LogError("API", "ID conv error", err)
	}
	res, _ := middlewares.GetSession(c)
	if id == res.UserID {
		code := utils.GenerateRandomInt(5)
		c.BindJSON(&user)
		errs := models.SendVerifyPhoneUser(&user, id, code, medium)
		if errs != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			plugins.LogError("API", "Error sending verification code to the user's phone", errs)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message":    "Phone verification code sent successfully",
				"statusCode": 200,
			})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "Unauthorized to access resource",
			"statusCode": 400,
		})
	}
}

//VerifyPhoneCodeUser - Verifies User's phone code
// @Summary Verifies user's phone verification code.
// @Description Sets phone verified to true if the code is correct and not expired.
// @Tags User Auth
// @Accept  json
// @Produce json
// @Success 200 {object} models.User
// @Router /user/verify-phone-code/:id [post]
func VerifyPhoneCodeUser(c *gin.Context) {
	var user models.User
	idString := c.Params.ByName("id")
	token := c.Query("token")
	id, err := utils.ConvertStringToInt(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":    "Id conv error",
			"statusCode": 500,
		})
		plugins.LogError("API", "ID conv error", err)
	}
	res, _ := middlewares.GetSession(c)
	if id == res.UserID {
		c.BindJSON(&user)
		state, errs := models.VerifyPhoneUser(&user, id, token)
		if errs != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			plugins.LogError("API", "An error occured while verifying user's phone number", errs)
		} else {
			if state == true && errs == nil {
				c.JSON(http.StatusOK, gin.H{
					"message":    "Phone verifed successfully",
					"statusCode": 200,
				})
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"message":    "Token expired or invalid",
					"statusCode": 400,
				})
			}
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "An error",
			"statusCode": 400,
		})
	}
}

//UpdateUser - Update an existing User
// @Summary Updates user based on given ID
// @Tags User Auth
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /user/{id} [patch]
// @Security ApiKeyAuth
func UpdateUser(c *gin.Context) {
	var user models.User
	idString := c.Params.ByName("id")
	id, err := utils.ConvertStringToInt(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":    "Id conv error",
			"statusCode": 500,
		})
		plugins.LogError("API", "ID conv error", err)
	}
	res, _ := middlewares.GetSession(c)
	if res.IsAdmin || id == res.UserID {
		errs := models.GetUser(&user, id)
		if errs != nil {
			c.JSON(http.StatusNotFound, user)
			plugins.LogError("API", "Error User not found", errs)
		}
		c.BindJSON(&user)
		err = models.UpdateUser(&user, id)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			plugins.LogError("API", "Error updating user", err)
		} else {
			c.JSON(http.StatusOK, user)
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message":    "An error",
			"statusCode": 400,
		})
	}
}

//DeleteUser - Deletes User
// @Summary Deletes a user based on given ID
// @Tags User Auth
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /user/{id} [delete]
// @Security ApiKeyAuth
func DeleteUser(c *gin.Context) {
	idString := c.Params.ByName("id")
	id, err := utils.ConvertStringToInt(idString)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":    "Id conv error",
			"statusCode": 500,
		})
		plugins.LogError("API", "Offset conv error", err)
	}
	errs := models.DeleteUser(id)
	if errs != nil {
		c.AbortWithStatus(http.StatusNotFound)
		plugins.LogError("API", "An error occured while deleting user", errs)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":    "Deleted successfully",
			"statusCode": 200,
		})
	}
}
