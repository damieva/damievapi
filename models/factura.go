package models

import "time"

type Factura struct {
	FacturaID uint             `gorm:"primaryKey" json:"facturaid"`
	PersonaID uint             `json:"personaid"`
	Fecha     time.Time        `json:"fecha"`
	Pais      string           `gorm:"size:200" json:"pais"`
	Ciudad    string           `gorm:"size:200" json:"ciudad"`
	Detalles  []FacturaDetalle `gorm:"foreignKey:FacturaID" json:"detalles"` // de esta forma definimos la clave foranea de factura en la tabla factura detalle
}
