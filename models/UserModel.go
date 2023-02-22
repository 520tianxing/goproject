package models

import "github.com/gin-gonic/gin"

type UserModel struct {
	Model
	Name string
	Age  string
}

func NewUserModel() *UserModel {
	return &UserModel{
		Model: Model{},
	}
}

func (u *UserModel) M(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": '3',
		"age":  "2",
	})
}
