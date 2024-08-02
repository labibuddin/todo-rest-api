package routes

import (
	"todoAPI/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRouter(app *fiber.App)  {
	app.Get("/todos", handlers.GetTodos)
    app.Get("/todos/:id", handlers.GetTodo)
    app.Post("/todos", handlers.CreateTodo)
    app.Put("/todos/:id", handlers.UpdateTodo)
    app.Delete("/todos/:id", handlers.DeleteTodo)
}