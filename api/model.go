package main

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint           `json:"id" gorm:"id,primarykey"`
	CreatedAt time.Time      `json:"createdAt" gorm:"created_at"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"deleted_at,index"`
}
