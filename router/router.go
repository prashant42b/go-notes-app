package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	noteRoutes "github.com/prashant42b/go-notes-app/internal/routes/note"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api", logger.New())

	//Setup node routers
	noteRoutes.SetupNoteRoutes(api)

}
