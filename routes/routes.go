package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wpcodevo/golang-fiber-mysql/controllers"
)

func RegisterRoutes(app *fiber.App) {

	notesAPI := app.Group("/notes")
	notesAPI.Get("", controllers.FindNotes)
	notesAPI.Delete("", controllers.DeleteNote)
	notesAPI.Post("", controllers.CreateNoteHandler)
	notesAPI.Patch("", controllers.UpdateNote)
	notesAPI.Get("/:noteId", controllers.FindNoteById)

}
