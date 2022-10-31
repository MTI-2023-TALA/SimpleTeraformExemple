package main

import (
	"os"
	"simple-teraform-exemple/apps/api/controller"
	"simple-teraform-exemple/apps/api/initializers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
    AllowOrigins: "*",
    AllowHeaders:  "Origin, Content-Type, Accept",
	}))

	app.Post("/todo", controller.CreateTodo)
	app.Put("/todo/:id", controller.UpdateTodo)
	app.Get("/todo", controller.GetAllTodo)
	app.Get("/todo/:id", controller.GetOneTodo)
	app.Delete("/todo/:id", controller.DeleteTodo)

	port := os.Getenv("PORT")
	if (len(port) == 0) {
		port = ":3000"
	}
	app.Listen(port)
}
