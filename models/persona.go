package models

type Persona struct {
	// 'json:' Definimos el campo que se estará usando al momento de convertir este objeto a JSON, para las funciones JSON y BodyParser seran importantes estas tags
	PersonaID uint      `gorm:"primaryKey" json:"personaid"`
	Email     string    `gorm:"unique;size:200" json:"email"`
	Password  string    `gorm:"size:200" json:"password"`
	Nombre    string    `gorm:"size:200" json:"nombre"`
	Facturas  []Factura `gorm:"foreignKey:PersonaID" json:"facturas"` // de esta forma definimos la clave foranea de persona en la tabla facturas
}

func (p Persona) ObtenerUsuarios() ([]Persona, error) {
	var personas []Persona
	err := DB.Find(&personas).Error

	return personas, err
}

func (p *Persona) RegistrarUsuario() error {
	return DB.Create(p).Error
}

func (p *Persona) Login() (bool, error) {
	var count int64
	result := DB.Model(&Persona{}).Where("email = ? AND password = ?", p.Email, p.Password).First(&p).Count(&count) // busqueda en la tabla personas para que email y pass sea la que le estamos pasando, una vez que obtenga la fila guardamos ese registro en la vble persona
	return count > 0, result.Error
}
