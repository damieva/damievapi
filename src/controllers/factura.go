package controllers

import (
	"DamievAPI/models"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type FacturaController struct{}

func NewFacturaController() *FacturaController {
	return &FacturaController{}
}

func (f FacturaController) ConfigPath(app *fiber.App) *fiber.App {
	app.Get("/", f.HandlerListarFacturas)
	app.Post("/", f.HandlerRegistrarFactura)
	return app
}

func (f FacturaController) HandlerRegistrarFactura(c *fiber.Ctx) error {
	var factura models.Factura
	if err := c.BodyParser(&factura); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Mensaje: err.Error()})
	}

	if err := factura.RegistrarFactura(); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Mensaje: err.Error()})
	}

	for i, v := range factura.Detalles {
		fmt.Println(i, v)
	}

	return c.JSON(factura)
}

func (f FacturaController) HandlerListarFacturas(c *fiber.Ctx) error {
	facturas, err := models.Factura{}.ObtenerFacturas()

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Mensaje: err.Error()})
	}

	return c.JSON(facturas)
}
