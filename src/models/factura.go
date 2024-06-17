package models

import (
	"time"

	"gorm.io/gorm/clause"
)

type Factura struct {
	FacturaID uint             `gorm:"primaryKey" json:"facturaid"`
	PersonaID uint             `json:"personaid"`
	Fecha     time.Time        `json:"fecha"`
	Pais      string           `gorm:"size:200" json:"pais"`
	Ciudad    string           `gorm:"size:200" json:"ciudad"`
	Detalles  []FacturaDetalle `gorm:"foreignKey:FacturaID" json:"detalles"` // de esta forma definimos la clave foranea de factura en la tabla factura detalle
}

func (f *Factura) RegistrarFactura() error {
	return DB.Create(f).Error
}

func (f Factura) ObtenerFacturas() ([]Factura, error) {
	var facturas []Factura
	// Deberíamos cargar las tablas detalles, productos y su asociacion en la query pero nuestro modelo de datos no lo permite
	err := DB.Preload(clause.Associations).Preload("Detalles").Find(&facturas).Error // Cargamos las relacions de la tabla detalle con el preload

	return facturas, err
}

func (f *Factura) ObtenerFacturaPorID(id uint) error {
	f.FacturaID = id
	return DB.Preload(clause.Associations).Preload("Detalles").First(&f).Error
}

func (f Factura) ObtenerFacturaPorUsuario(idusuario uint) ([]Factura, error) {
	var facturas []Factura
	err := DB.Where("persona_id = ?", idusuario).Preload(clause.Associations).Preload("Detalles").First(&facturas).Error
	return facturas, err
}
