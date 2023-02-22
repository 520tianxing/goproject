package controllers

import (
	"goproject/src/go_code/ginproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Contoller
	UserModel models.UserModel
}

func NewUserController() *UserController {
	return &UserController{
		Contoller: Contoller{},
		UserModel: *models.NewUserModel(),
	}
}

func (u *UserController) U(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "wohao",
	})
}

func (u *UserController) HtmlU(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}
