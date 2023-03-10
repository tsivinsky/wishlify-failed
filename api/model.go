package main

import (
	"time"
)

type Model struct {
	ID        uint       `json:"id" gorm:"id,primarykey"`
	CreatedAt time.Time  `json:"createdAt" gorm:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" gorm:"deleted_at,index"`
}
