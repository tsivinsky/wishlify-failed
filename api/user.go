package main

import "time"

type User struct {
	Model

	Email       string     `json:"email" gorm:"email,unique"`
	Username    string     `json:"username" gorm:"username,unique"`
	Password    string     `json:"-" gorm:"password"`
	ConfirmedAt *time.Time `json:"confirmedAt" gorm:"confirmed_at"`
}
