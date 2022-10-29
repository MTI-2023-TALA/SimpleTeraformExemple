package controller

import (
	"simple-teraform-exemple/apps/api/initializers"
	"simple-teraform-exemple/apps/api/model"

	"github.com/gofiber/fiber/v2"
)

func CreateTodo(c *fiber.Ctx) error {
	// Get the request body
	var body struct {
		Title string
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}


	// Create a post
	todo := model.Todo{Title: body.Title}
	result := initializers.DB.Create(&todo)

	if result.Error != nil {
		return c.Status(500).SendString(result.Error.Error())
	}

	return c.JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	// Get the request body
	var body struct {
		Title string
	}

	if err := c.BodyParser(&body); err != nil {
		return err
	}

	// Update
	var todo model.Todo
	initializers.DB.First(&todo, c.Params("id"))
	initializers.DB.Model(&todo).Updates(model.Todo{Title: body.Title})
	return c.JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	// Delete
	var todo model.Todo
	initializers.DB.First(&todo, c.Params("id"))
	initializers.DB.Delete(&todo)

	return c.SendStatus(200)
}

func GetAllTodo(c *fiber.Ctx) error {
	// Get
	var todos []model.Todo
	initializers.DB.Find(&todos)

	return c.JSON(todos)
}

func GetOneTodo(c *fiber.Ctx) error {
	// Get
	var todo model.Todo
	initializers.DB.First(&todo, c.Params("id"))

	return c.JSON(todo)
}
