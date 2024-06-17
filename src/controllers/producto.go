package controllers

import (
	"DamievAPI/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ProductoController struct{}

// Constructor
func NewProductoController() *ProductoController {
	return &ProductoController{}
}

// Hereda de la interfaz microservicio:
func (p ProductoController) ConfigPath(app *fiber.App) *fiber.App {
	app.Get("/", p.HandlerObtenerProductos)
	app.Post("/", p.HandlerRegistrarProducto) // Registramos el path para registrar un producto
	return app
}

func (p ProductoController) HandlerObtenerProductos(c *fiber.Ctx) error {
	productos, err := models.Producto{}.ObtenerProductos()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Mensaje: err.Error()})
	}

	return c.JSON(productos)
}

func (p ProductoController) HandlerRegistrarProducto(c *fiber.Ctx) error {
	var producto models.Producto

	if err := c.BodyParser(&producto); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Mensaje: err.Error()})
	}

	if err := producto.Registrar(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Mensaje: err.Error()})
	}

	return c.JSON(producto) // El registro se hizo bien y devolvemos el producto actualizado
}
