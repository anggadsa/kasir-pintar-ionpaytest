package services

import (
	"kasir-pintar-ionpaytest/core"
	"kasir-pintar-ionpaytest/models"

	"github.com/jinzhu/gorm"
)

type Registration struct {
	core.Service
}

func (s *Registration) All(registration interface{}) (*gorm.DB, []*models.Registration) {
	var list []*models.Registration
	result := s.Db().Where(registration).Find(&list)
	return result, list
}

func (s *Registration) Create(registration *models.Registration) (*gorm.DB, *models.Registration) {
	result := s.Db().Create(&registration)

	// bytes, _ := json.Marshal(registration)
	// fmt.Println(string(bytes))
	return result, registration
}
