# General notes:
* Los nombres de las funciones en go tienen la siguiente connotación:
    - Minuscula: funciones privadas solo accesibles desde el propio package
    - Mayuscula: funciones publicas accesibles desde cualquier parte del codigo.
* Si un objeto implementa todos los metodos de una interfaz se pude interpretar directamente como herencia.
* A una vbla nueva se le asigna un valor con ':=', si usamos una variable ya creada le podemos asignar un valor simeplemente con '='

# DamievAPI
API written in Golang for my personal project

Fiber module: 
* Modulo para crear la api rest
* Installation: 'go get github.com/gofiber/fiber/v2'

Gorm module:
* Modulo para 
* Installation: go get -u gorm.io/gorm

Viper module:
* Modulo para utilizar archivos de conf, vbles de entorno, etc
* Installation: go get github.com/spf13/viper

JWT module:
* go get -u github.com/golang-jwt/jwt