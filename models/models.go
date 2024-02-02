package models

type Register struct {
	Username		string	`json:"username" binding:"required"`
	Password		string	`json:"password" binding:"required"`
	Fullname		string	`json:"fullname" binding:"required"`
	Avatar			string	`json:"avatar" binding:"required"`
}

type Login struct {
	Username 		string 	`json:"username"`
	Password 		string 	`json:"password"`
}