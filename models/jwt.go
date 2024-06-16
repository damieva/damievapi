package models

import "github.com/golang-jwt/jwt"

type JWTClient struct {
	Usuario Persona `json:"usuario"`
	Token   string  `json:"token"`
}

type Claims struct {
	ID                 uint   `json:"id"`
	Email              string `json:"email"`
	jwt.StandardClaims        // esto es un objeto cuya estructura ya satisface los atributos requeridos por la interfaz claims
}
