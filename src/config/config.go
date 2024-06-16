package config

import "github.com/gofiber/fiber/v2"

// Definimos la interrfaz de los microservicios
type MicroServicio interface { //

	ConfigPath(app *fiber.App) *fiber.App
}

// Esta funcion Use nos servirá para añadir cada microservicio a la api, podria ir tmabién en main
func Use(prefix string, r fiber.Router, micro MicroServicio) {

	r.Mount(prefix, micro.ConfigPath(fiber.New()))
}
