package controllers

import (
	"net/http"
	"time"

	"github.com/DamievaIT/DamievAPI-2/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// En cada controlador indicaremos toda la logica de cada microservicio
// Definicion de un microservicio que va a implementar la interfaz definida en config.go
type UsuarioController struct{}

// Constructor:
func NewUsuarioController() *UsuarioController {
	return &UsuarioController{}
}

var secret = []byte("carrito_compra") // Secreto para firmar tokens

// A continuacion definimos los metodos del UsuarioController, vemos que son metodos porque tiene el objeto delante
// Al definir el método ConfigPath podemos entender que Usuario controller hereda de la interfaz microservicio
func (u UsuarioController) ConfigPath(app *fiber.App) *fiber.App {
	app.Get("/", u.ObtenerUsuario)    //Registramos el path de obtener todos los uaurios
	app.Post("/", u.RegistrarUsuario) // Registramos el path para registrar un usuario
	app.Post("/login", u.HandlerLogin)
	return app
}

func (u UsuarioController) ObtenerUsuario(c *fiber.Ctx) error {
	usuarios, err := models.Persona{}.ObtenerUsuarios()

	if err != nil {
		// Retornamos un codigo de error http y ademas un JSON con el mensaje de error que nos ha dado el programa
		return c.Status(http.StatusBadRequest).JSON(models.Error{Mensaje: err.Error()})
	}

	// Al momento de convertir la vble usuario a JSON va a utilizar las tags
	return c.JSON(usuarios)
}

func (u UsuarioController) RegistrarUsuario(c *fiber.Ctx) error {
	var usuario models.Persona
	err := c.BodyParser(&usuario) // Pasamos la direccion de memoria del objeto usuario y la funcion BodyParser directamente parsea los valores recibidos en el body a los campos del objeto Persona

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Mensaje: err.Error()})
	}

	err = usuario.RegistrarUsuario() // Llamamos al metodos Registar usuario con un objeto de tipo Persona que ya tendrá todos los valores para esa persona
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Mensaje: err.Error()})
	}

	return c.JSON(usuario)
}

func (u UsuarioController) HandlerLogin(c *fiber.Ctx) error {
	var usuario models.Persona

	if err := c.BodyParser(&usuario); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Mensaje: err.Error()})
	}

	ok, err := usuario.Login()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Mensaje: err.Error()})
	}

	if !ok {
		c.Status(http.StatusBadRequest).JSON(models.Error{Mensaje: "El usuario o pass no son correctos"})
	}

	tokenString, err := GenerarToken(usuario)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.Error{Mensaje: err.Error()})
	}

	respuesta := models.JWTClient{
		Usuario: usuario,
		Token:   tokenString,
	}

	return c.JSON(respuesta)
}

func GenerarToken(usuario models.Persona) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{ //creamos un nuevo token y definimos sus claims
		"id":    usuario.PersonaID,
		"email": usuario.Email,
		"exp":   time.Now().Add(time.Hour * 8).Unix(), // el token va a expirar en 8 horas y además la fecha tendra el formato Unix
	})

	return token.SignedString(secret) // Retornamos el token firmado
}
