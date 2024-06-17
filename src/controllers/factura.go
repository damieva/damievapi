package controllers

import (
	"DamievAPI/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type FacturaController struct{}

func NewFacturaController() *FacturaController {
	return &FacturaController{}
}

func (f FacturaController) ConfigPath(app *fiber.App) *fiber.App {
	app.Get("/", f.HandlerListarFacturas)
	app.Post("/", f.HandlerRegistrarFactura)
	app.Get("/:id", f.HandlerObtenerFacturaPorID) // mandamos el parametro en la url
	app.Get("/usuario/:id", f.HandlerObtenerFacturaPorUsuarioID)
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

func (f FacturaController) HandlerObtenerFacturaPorID(c *fiber.Ctx) error {
	// Obtenemos el paramtro ID de la peticion
	idstring := c.Params("id")
	id, err := strconv.Atoi(idstring)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Mensaje: err.Error()})
	}

	var factura models.Factura

	factura.ObtenerFacturaPorID(uint(id))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Mensaje: err.Error()})
	}

	return c.JSON(factura)
}

func (f FacturaController) HandlerObtenerFacturaPorUsuarioID(c *fiber.Ctx) error {
	idstring := c.Params("id")
	id, err := strconv.Atoi(idstring)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Mensaje: err.Error()})
	}

	facturas, err := models.Factura{}.ObtenerFacturaPorUsuario(uint(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Mensaje: err.Error()})
	}

	return c.JSON(facturas)
}
