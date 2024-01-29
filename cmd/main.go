package main

import (
	"SampleCRM/config"
	"SampleCRM/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	r := gin.Default()
	controllers.RegisterBookRoutes(r)
	r.Run() // Graceful Shutdown?
}
