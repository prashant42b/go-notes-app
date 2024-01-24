package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prashant42b/crud-task/config"
	"github.com/prashant42b/crud-task/database"
	"github.com/prashant42b/crud-task/routes"
	util "github.com/prashant42b/crud-task/utils"
)

func main() {
	//importing env
	util.ImportENV()
	config.LoadConfig()

	//creating new fiber app
	app := fiber.New()

	//establish connection to postgres db
	database.ConnectDB()

	//redirecting app to its routes
	routes.SetupRoutes(app)

	//app listens on port 3000
	app.Listen(":3000")

}
