package core

import (
	"time"

	"kasir-pintar-ionpaytest/config"

	"github.com/jinzhu/gorm"
)

type Model struct {
	ID        int        `json:"-" gorm:"primarykey"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

func (m *Model) Db() *gorm.DB {
	return config.Db()
}
