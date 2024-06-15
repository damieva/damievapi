package models

import (
	"github.com/DamievaIT/DamievAPI-2/db"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	DB = db.GetConn()

	// Cuando le pasamos las estucturas a la funcion automigrate ya automaticamente nos crea la estructura de la DB
	DB.AutoMigrate(&Persona{})
	DB.AutoMigrate(&Producto{})
	DB.AutoMigrate(&Factura{})
	DB.AutoMigrate(&FacturaDetalle{})
}
