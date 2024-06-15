package models

type FacturaDetalle struct {
	FacturaID  uint    `gorm:"primaryKey" json:"facturaid"` // de esta manera hacemos que las claves foraneas sean a la vez primarias, tipico de una tabla renacida
	ProductoID uint    `gorm:"primaryKey" json:"productoid"`
	Cantidad   float32 `json:"cantidad"`
	Precio     float32 `json:"precio"`
	Descuento  float32 `json:"descuento"`
}
