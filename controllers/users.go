package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/JcsnP/go-jwt-app/config/db"
	"github.com/JcsnP/go-jwt-app/config/schema"
	"github.com/JcsnP/go-jwt-app/utils"
)

// find user by id
func FindUserById(c *gin.Context) {
	header := c.Request.Header.Get("Authorization")
	tokenString := strings.Replace(header, "Bearer ", "", 1)

	token, err := utils.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
	}

	id := c.Param("id")
	var user schema.User

	// find user
	db.DB.Where("id = ?", id).First(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": user,
		"accessToken": token,
	})
}

// get all users
func GetAllUsers(c *gin.Context) {
	var users []schema.User

	// get all users
	if err := db.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": users,
	})
}

// delete user
func DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	var user schema.User

	// find the user
	if err := db.DB.Find(&user, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// then delete the user
	if err := db.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}