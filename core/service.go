package core

import (
	"kasir-pintar-ionpaytest/config"

	"github.com/jinzhu/gorm"
)

type Service struct {
}

func (m *Service) Db() *gorm.DB {
	return config.Db()
}
