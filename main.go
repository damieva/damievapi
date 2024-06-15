package main

import (
	"github.com/DamievaIT/DamievAPI-2/config"
	"github.com/DamievaIT/DamievAPI-2/controllers"
	"github.com/DamievaIT/DamievAPI-2/models"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// Creamos la api
	app := fiber.New()

	models.InitDB()
	api := app.Group("/api/v1")

	// Añadimos los microservicios a la api
	config.Use("/usuarios", api, controllers.NewUsuarioController())

	// Dejamos a la api escuchando en el puerto 3000
	app.Listen(":3000")
}
