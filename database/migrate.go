package database

import (
	"kasir-pintar-ionpaytest/config"
	"kasir-pintar-ionpaytest/models"
	"log"
)

func Migrate() {
	log.Printf("Migrate: Start")
	db := config.Db()
	db.AutoMigrate(
		&models.Registration{},
	)
	log.Printf("Migrate: Success")
}
