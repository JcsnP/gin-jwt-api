package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/JcsnP/go-jwt-app/config/db"
	"github.com/JcsnP/go-jwt-app/config/schema"
	"github.com/JcsnP/go-jwt-app/controllers"
)



func main() {
	gin.DisableConsoleColor()

	f, _ := os.Create(fmt.Sprintf("./logs/%s.log", time.Now().Format("01-02-2001_15:04:05")))
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	r := gin.Default()
	db.CreateDatabase()
	
	// delete user's record before do something
	// utils.RemoveRecord()

	db.DB.AutoMigrate(&schema.User{})

	// users
	r.GET("/users", controllers.GetAllUsers)
	r.GET("/users/:id", controllers.FindUserById)
	r.DELETE("/users/:id", controllers.DeleteUserByID)

	// auth
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/me", controllers.GetMe)

	r.Run()
}  