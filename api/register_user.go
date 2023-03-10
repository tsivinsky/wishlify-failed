package main

import "github.com/gofiber/fiber/v2"

type RegisterUserBody struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleRegisterUser(c *fiber.Ctx) error {
	var body RegisterUserBody

	if err := c.BodyParser(&body); err != nil {
		return MakeApiError(400, err.Error())
	}

	hashedPassword, err := HashPassword(body.Password)
	if err != nil {
		return MakeApiError(500, err.Error())
	}

	if tx := Db.Where("email = ?", body.Email).First(&User{}); tx.RowsAffected > 0 {
		return MakeApiValidationError("email", "Email already registered")
	}

	if tx := Db.Where("username = ?", body.Username).First(&User{}); tx.RowsAffected > 0 {
		return MakeApiValidationError("username", "Username already taken")
	}

	user := &User{
		Email:    body.Email,
		Username: body.Username,
		Password: hashedPassword,
	}

	tx := Db.Create(user)
	if tx.Error != nil {
		return MakeApiError(500, tx.Error.Error())
	}

	// TODO: generate jwt tokens here and send them in response

	return c.Status(201).JSON(user)
}
