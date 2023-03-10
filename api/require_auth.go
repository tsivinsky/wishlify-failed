package main

import "github.com/gofiber/fiber/v2"

func HandleRequireAuth(c *fiber.Ctx) error {
	accessToken, err := GetAccessTokenFromHeader(c)
	if err != nil {
		return MakeApiError(401, err.Error())
	}

	userId, err := ValidateAccessToken(accessToken)
	if err != nil {
		return MakeApiError(401, err.Error())
	}

	var user User
	tx := Db.Where("id = ?", userId).First(&user)
	if tx.Error != nil {
		return MakeApiError(401, err.Error())
	}

	c.Locals("userId", userId)

	return c.Next()
}
