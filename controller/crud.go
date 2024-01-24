package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/prashant42b/crud-task/database"
	"github.com/prashant42b/crud-task/model"
)

func CreateRecord(c *fiber.Ctx) error {
	db := database.DB
	todo := new(model.Todo)

	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Check your input", "data": err})
	}

	todo.ID = uuid.New()

	err = db.Create(&todo).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create Data", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Record Created", "data": todo})

}

func GetRecord(c *fiber.Ctx) error {
	db := database.DB
	var todo model.Todo

	id := c.Params("todoId")
	db.Find(&todo, "id = ?", id)

	if todo.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not find Note", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Note found", "data": todo})

}
func UpdateRecord(c *fiber.Ctx) error {
	db := database.DB
	var todo model.Todo

	type updateDB struct {
		Title    string `json:"title"`
		SubTitle string `json:"sub_title"`
		Text     string `json:"Text"`
	}

	id := c.Params("todoId")
	db.Find(&todo, "id = ?", id)

	if todo.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not find note", "data": nil})
	}

	var updateData updateDB
	err := c.BodyParser(&updateData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	todo.Title = updateData.Title
	todo.SubTitle = updateData.SubTitle
	todo.Text = updateData.Text

	db.Save(&todo)

	return c.JSON(fiber.Map{"status": "success", "message": "Note found and updated", "data": todo})

}
func DeleteRecord(c *fiber.Ctx) error {
	db := database.DB
	var todo model.Todo

	id := c.Params("todoId")
	db.Find(&todo, "id = ?", id)

	if todo.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not find Data", "data": nil})
	}

	err := db.Delete(&todo, "id = ?", id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Unable to delete Data", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Note deleted"})

}
