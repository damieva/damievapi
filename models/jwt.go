package models

type JWTClient struct {
	Usuario Persona `json:"usuario"`
	Token   string  `json:"token"`
}
