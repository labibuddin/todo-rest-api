package main

import (
	"todoAPI/database"
	"todoAPI/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    database.ConnectDB()

	routes.SetupRouter(app)
    

    app.Listen(":3000")
}
