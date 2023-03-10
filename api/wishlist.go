package main

import (
	"regexp"
	"strings"
)

type Wishlist struct {
	Model

	Name        string    `json:"name" gorm:"name"`
	Description string    `json:"description" gorm:"description"`
	DisplayName string    `json:"displayName" gorm:"display_name"`
	UserId      uint      `json:"userId" gorm:"user_id"`
	User        User      `json:"user" gorm:"user"`
	Products    []Product `json:"products" gorm:"products"`
}

var re = regexp.MustCompile(`[^A-z|0-9|А-я]+`)

func GenerateWishlistDisplayName(name string) string {
	lowerCased := strings.ToLower(name)

	displayName := re.ReplaceAll([]byte(lowerCased), []byte("-"))

	return string(displayName)
}
