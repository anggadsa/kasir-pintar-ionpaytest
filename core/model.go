package core

import (
	"time"

	"kasir-pintar-ionpaytest/config"

	"github.com/jinzhu/gorm"
)

type Model struct {
	ID        int        `json:"id" gorm:"primarykey"`
	CreatedAt time.Time  `json:"createdAt,omitempty"`
	UpdatedAt time.Time  `json:"updatedAt,omitempty"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`
}

func (m *Model) Db() *gorm.DB {
	return config.Db()
}
