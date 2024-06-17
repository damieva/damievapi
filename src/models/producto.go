package models

type Producto struct {
	ProductoID     uint             `gorm:"primaryKey" json:"productoid"`
	NombreProducto string           `gorm:"size:200" json:"nombreProducto"`
	Descripcion    string           `gorm:"size:500" json:"descripcion"`
	PrecioUnitario float32          `json:"precioUnitario"`
	CantidadStock  float32          `json:"cantidadStock"`
	Descontinuado  bool             `json:"descontinuado"`
	Detalles       []FacturaDetalle `gorm:"foreignKey:ProductoID" json:"detalles"` // De esta forma definimos la clave foranea de producto en la tabla factura detalle
}

func (p Producto) ObtenerProductos() ([]Producto, error) {
	var productos []Producto

	err := DB.Find(&productos).Error
	return productos, err
}

func (p *Producto) Registrar() error { //Al enviar la direccion de memoria (p *Producto) del producto podremos ver los campos (que el cliente no intruducirá) después de ejecutar el registro (como por ejemplo ProductoID)
	return DB.Create(p).Error
}
