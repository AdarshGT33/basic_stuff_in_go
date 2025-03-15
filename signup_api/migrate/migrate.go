package main

import (
	"main.go/initializers"
	"main.go/models"
)

func init() {
	initializers.LoadEnvs()
	initializers.CreateDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
