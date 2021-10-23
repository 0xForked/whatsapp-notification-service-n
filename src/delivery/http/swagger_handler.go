package http

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func NewSwaggerHandler(app *fiber.App) {
	// Routes for GET method:

	//default
	app.Get("/swagger/*", swagger.Handler)
}
