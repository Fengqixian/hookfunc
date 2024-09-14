package model

import (
	"gorm.io/gorm"
	"time"
)

type Resource struct {
	Id           uint   `gorm:"primarykey"`
	ResourceLink string `gorm:"unique;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (m *Resource) TableName() string {
	return "resource"
}
