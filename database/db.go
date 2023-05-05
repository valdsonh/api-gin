package database

import (
	"log"
	"valdson/api-gin/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	connStr := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connStr))

	if err != nil {
		log.Panic("Erro ao conectar o banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
