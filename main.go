package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prashant42b/go-notes-app/config"
	"github.com/prashant42b/go-notes-app/database"
	"github.com/prashant42b/go-notes-app/router"
	"github.com/prashant42b/go-notes-app/util"
)

func main() {
	//start a new fiber app
	util.ImportENV()
	config.LoadConfig()

	app := fiber.New()
	database.ConnectDB()
	router.SetupRoutes(app)

	//port no 3000
	app.Listen(":3000")
}
