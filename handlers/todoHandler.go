package handlers

import (
	"todoAPI/database"
	"todoAPI/models"

	"github.com/gofiber/fiber/v2"
)

// GetTodos - Get all todos
func GetTodos(c *fiber.Ctx) error {
    var todos []models.Todo
    database.DB.Db.Find(&todos)
    return c.Status(200).JSON(todos)
}

// GetTodo - Get single todo by ID
func GetTodo(c *fiber.Ctx) error {
    id := c.Params("id")
    var todo models.Todo
    result := database.DB.Db.First(&todo, id)

    if result.Error != nil {
        return c.Status(404).JSON(fiber.Map{"message": "Todo not found"})
    }

    return c.Status(200).JSON(todo)
}

// CreateTodo - Create new todo
func CreateTodo(c *fiber.Ctx) error {
    todo := new(models.Todo)

    if err := c.BodyParser(todo); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": "Failed to parse request body"})
    }

    database.DB.Db.Create(&todo)
    return c.Status(201).JSON(todo)
}

// UpdateTodo - Update existing todo
func UpdateTodo(c *fiber.Ctx) error {
    id := c.Params("id")
    var todo models.Todo
    result := database.DB.Db.First(&todo, id)

    if result.Error != nil {
        return c.Status(404).JSON(fiber.Map{"message": "Todo not found"})
    }

    if err := c.BodyParser(&todo); err != nil {
        return c.Status(400).JSON(fiber.Map{"message": "Failed to parse request body"})
    }

    database.DB.Db.Save(&todo)
    return c.Status(200).JSON(todo)
}

// DeleteTodo - Delete a todo by ID
func DeleteTodo(c *fiber.Ctx) error {
    id := c.Params("id")
    var todo models.Todo
    result := database.DB.Db.First(&todo, id)

    if result.Error != nil {
        return c.Status(404).JSON(fiber.Map{"message": "Todo not found"})
    }

    database.DB.Db.Delete(&todo)
    return c.Status(204).SendString("Todo successfully deleted")
}
