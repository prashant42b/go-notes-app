package noteHandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/prashant42b/go-notes-app/database"
	"github.com/prashant42b/go-notes-app/internal/model"
)

func GetNotes(c *fiber.Ctx) error {
	db := database.DB
	var notes []model.Note
	db.Find(&notes)

	if len(notes) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No notes found", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Notes found", "data": notes})

}

func CreateNotes(c *fiber.Ctx) error {
	db := database.DB
	note := new(model.Note)

	err := c.BodyParser(note)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Check your input", "data": err})
	}

	note.ID = uuid.New()

	err = db.Create(&note).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create Note", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Note Created", "data": note})

}

func GetNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	id := c.Params("noteId")
	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not find Note", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Note found", "data": note})

}
func UpdateNote(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	type updateNote struct {
		Title    string `json:"title"`
		SubTitle string `json:"sub_title"`
		Text     string `json:"Text"`
	}

	id := c.Params("noteId")
	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not find note", "data": nil})
	}

	var updateNoteData updateNote
	err := c.BodyParser(&updateNoteData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	note.Title = updateNoteData.Title
	note.SubTitle = updateNoteData.SubTitle
	note.Text = updateNoteData.Text

	db.Save(&note)

	return c.JSON(fiber.Map{"status": "success", "message": "Note found and updated", "data": note})

}
func DeleteNode(c *fiber.Ctx) error {
	db := database.DB
	var note model.Note

	id := c.Params("noteId")
	db.Find(&note, "id = ?", id)

	if note.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Could not find Note", "data": nil})
	}

	err := db.Delete(&note, "id = ?", id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Unable to delete Note", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Note deleted"})

}
