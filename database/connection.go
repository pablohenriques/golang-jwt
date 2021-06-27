package database

import (
	"golang_jwt/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect()  {
	dsn := "host=localhost user=postgres password=docker dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Não foi possível conecta-se ao banco de dados.")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
