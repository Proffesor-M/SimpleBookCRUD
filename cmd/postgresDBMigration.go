package main

import (
	"SampleCRM/config"
	"SampleCRM/models"
	"fmt"
)

func init() {
	config.LoadEnvVariables()
	config.ConnectDB()
}

func main() {
	config.DB.AutoMigrate(&models.Book{})
	fmt.Println("Done")
}
