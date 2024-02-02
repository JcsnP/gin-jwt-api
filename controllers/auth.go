package controllers

import (
	"net/http"
	"strings"

	"github.com/JcsnP/go-jwt-app/config/db"
	"github.com/JcsnP/go-jwt-app/config/schema"
	"github.com/JcsnP/go-jwt-app/models"
	"github.com/JcsnP/go-jwt-app/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var json models.Register
	var userExists schema.User

	// check if user already exists in database
	db.DB.Where("username = ?", json.Username).First(&userExists)
	if userExists.ID > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "username already exists",
		})
		return
	}

	// create user
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), bcrypt.DefaultCost)
	user := schema.User{Username: json.Username, Password: string(hashedPassword), Fullname: json.Fullname, Avatar: json.Avatar}

	db.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": user,
	}) 	
}

func Login(c *gin.Context) {
	var json models.Login

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// check if user exists
	var userExists schema.User
	db.DB.Where("username = ?", json.Username).First(&userExists)
	if userExists.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
		return
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(json.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	token := utils.GenerateToken(userExists.ID)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"accessToken": token,
	})
}

func GetMe(c *gin.Context) {
	header := c.Request.Header.Get("Authorization")
	accessToken := strings.Replace(header, "Bearer ", "", 1)

	id, err := utils.ValidateToken(accessToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	var user schema.User
	db.DB.Where("id = ?", id).First(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data": user,
	})
}