package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetConn() *gorm.DB {
	//dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", pghostname, pguser, pgpassword, pgdb, pgport)
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err) // si hay algun error para la ejecucion del programa
	}
	return db
}
