package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDatabase() (err error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", Env.DBHost, Env.DBUser, Env.DBPassword, Env.DBName)

	Db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		return
	}

	err = Db.AutoMigrate(&User{}, &Wishlist{}, &Product{})

	return
}
