package routes

import (
	"github.com/akers1023/Smart-Waste-Management-System/internal/app/handlers"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App, uh *handlers.UserHandler) {

	app.Post("/users/register", uh.Register)
	app.Get("/users/:id", uh.ViewInfomationUser)
}
