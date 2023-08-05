package database

import (
	"kasir-pintar-ionpaytest/database/seeds"
	"log"
)

func Seeder() {
	log.Printf("Seed: Start")
	seeds.Registrations()
	log.Printf("Seed: Success")
}
