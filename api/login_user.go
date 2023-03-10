package main

import "github.com/gofiber/fiber/v2"

type LoginUserBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandleLoginUser(c *fiber.Ctx) error {
	var body LoginUserBody

	if err := c.BodyParser(&body); err != nil {
		return MakeApiError(400, err.Error())
	}

	var user User
	if tx := Db.Where("email = ?", body.Email).First(&user); tx.RowsAffected == 0 {
		return MakeApiValidationError("email", "Incorrect email")
	}

	if ok := CheckPassword(user.Password, body.Password); !ok {
		return MakeApiValidationError("password", "Incorrect password")
	}

	return c.JSON(user)
}
