package main

import (
	"github.com/fyvri/go-qris/api/routes"
	"github.com/fyvri/go-qris/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	gin := gin.Default()
	routes.Setup(env, gin)

	gin.Run(":" + env.Port)
}
