package main

import (
	"context"
	"log"
	"simple-teraform-exemple/apps/api/controller"
	"simple-teraform-exemple/apps/api/initializers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {
	log.Printf("Fiber cold start")
	initializers.ConnectToDB()


	var app *fiber.App
	app = fiber.New()

	app.Use(cors.New(cors.Config{
    AllowOrigins: "*",
    AllowHeaders:  "Origin, Content-Type, Accept",
	}))

	app.Post("/todo", controller.CreateTodo)
	app.Put("/todo/:id", controller.UpdateTodo)
	app.Get("/todo", controller.GetAllTodo)
	app.Get("/todo/:id", controller.GetOneTodo)
	app.Delete("/todo/:id", controller.DeleteTodo)

	fiberLambda = fiberadapter.New(app)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return fiberLambda.ProxyWithContext(ctx, req)
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(Handler)
}
