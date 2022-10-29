package main

import (
	"simple-teraform-exemple/apps/api/initializers"
	"simple-teraform-exemple/apps/api/model"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&model.Todo{})
}
