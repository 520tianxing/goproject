package router

import (
	"goproject/src/go_code/ginproject/controllers"

	"github.com/gin-gonic/gin"
)

var UserController *controllers.UserController

type Router struct {
	Router         *gin.Engine
	UserController *controllers.UserController
}

func NewRouter() *Router {
	return &Router{
		Router:         gin.Default(),
		UserController: controllers.NewUserController(),
	}
}

// 请求地址写这里 以及所调用的方法
func (c *Router) Start() {
	c.Router.LoadHTMLGlob("view/*")
	c.Router.GET("/ping", c.UserController.U)
	c.Router.GET("/ph", c.UserController.HtmlU)
}

func (c *Router) Run(addr ...string) {
	if addr == nil {
		c.Router.Run()
	} else {
		c.Router.Run(addr...)
	}
}
