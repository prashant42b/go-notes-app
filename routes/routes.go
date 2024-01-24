package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prashant42b/crud-task/controller"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api")

	todo := api.Group("/todo")

	todo.Post("/", controller.CreateRecord)
	todo.Get("/:todoId", controller.GetRecord)
	todo.Put("/:todoId", controller.UpdateRecord)
	todo.Delete("/:todoId", controller.DeleteRecord)

}
