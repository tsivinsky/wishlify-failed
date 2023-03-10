package main

import (
	"github.com/gofiber/fiber/v2"
)

func HandleGetUser(c *fiber.Ctx) error {
	userId := c.Locals("userId").(uint)

	var user User
	if tx := Db.Where("id = ?", userId).First(&user); tx.Error != nil {
		return MakeApiError(500, tx.Error.Error())
	}

	return c.JSON(user)
}
