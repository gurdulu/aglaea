package main

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/gurdulu/aglaea/controllers"
	"github.com/gurdulu/aglaea/inits"
	"github.com/gurdulu/aglaea/middlewares"
	"log"
	"net/http"
)

//go:embed pages/*
var pagesFS embed.FS

func init() {
	inits.LoadEnvs()
	inits.ConnectDB()
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./pages/**/*.html")

	r.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	r.POST("/auth/signup", controllers.CreateUser)
	r.POST("/auth/login", controllers.Login)
	r.GET("/user/profile", middlewares.CheckAuth, controllers.GetUserProfile)

	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
