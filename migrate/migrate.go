package main

import (
	"marcellinuselbert-be/initializers"
	"marcellinuselbert-be/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
