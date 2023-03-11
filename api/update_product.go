package main

import "github.com/gofiber/fiber/v2"

type UpdateProductBody struct {
	Title      *string `json:"title"`
	Image      *string `json:"image"`
	WishlistId *uint   `json:"wishlistId"`
}

func HandleUpdateProduct(c *fiber.Ctx) error {
	productId, err := c.ParamsInt("productId")
	if err != nil {
		return MakeApiError(400, err.Error())
	}

	var body UpdateProductBody

	if err := c.BodyParser(&body); err != nil {
		return MakeApiError(400, err.Error())
	}

	var product *Product
	if tx := Db.First(&product, productId); tx.Error != nil {
		return MakeApiError(404, tx.Error.Error())
	}

	if body.Title != nil {
		product.Title = *body.Title
	}

	if body.Image != nil {
		product.Image = body.Image
	}

	if body.WishlistId != nil {
		product.WishlistId = *body.WishlistId
	}

	if tx := Db.Save(&product); tx.Error != nil {
		return MakeApiError(500, tx.Error.Error())
	}

	return c.JSON(product)
}
