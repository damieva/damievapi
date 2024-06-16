package controllers

import "github.com/gofiber/fiber/v2"

type ProductoController struct{}

// Constructor
func NewProductoController() *ProductoController {
	return &ProductoController{}
}

// Hereda de la interfaz microservicio:
func (u ProductoController) ConfigPath(app *fiber.App) *fiber.App {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Productos")
	})
	return app
}
