package main

import "github.com/gofiber/fiber/v2"

type CreateWishlistBody struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func HandleCreateWishlist(c *fiber.Ctx) error {
	var body CreateWishlistBody

	if err := c.BodyParser(&body); err != nil {
		return MakeApiError(400, err.Error())
	}

	if body.Name == "" {
		return MakeApiValidationError("name", "Name can't be empty")
	}

	userId := c.Locals("userId").(uint)

	var user User
	if tx := Db.Where("id = ?", userId).First(&user); tx.Error != nil {
		return MakeApiError(401, tx.Error.Error())
	}

	displayName := GenerateWishlistDisplayName(body.Name)

	if tx := Db.Where("display_name = ?", displayName).First(&Wishlist{}); tx.RowsAffected > 0 {
		return MakeApiValidationError("name", "Name already taken")
	}

	wishlist := Wishlist{
		Name:        body.Name,
		Description: body.Description,
		DisplayName: displayName,
		UserId:      userId,
		User:        user,
		Products:    []Product{},
	}

	if tx := Db.Create(&wishlist).Preload("User"); tx.Error != nil {
		return MakeApiError(500, tx.Error.Error())
	}

	return c.Status(201).JSON(wishlist)
}
