package main

import "github.com/gofiber/fiber/v2"

type AddProductBody struct {
	Title string  `json:"title"`
	Image *string `json:"image"`
}

func HandleAddProduct(c *fiber.Ctx) error {
	wishlistId, err := c.ParamsInt("id")
	if err != nil {
		return MakeApiError(400, err.Error())
	}

	var body AddProductBody

	if err := c.BodyParser(&body); err != nil {
		return MakeApiError(400, err.Error())
	}

	var wishlist *Wishlist
	if tx := Db.Preload("User").Preload("Products").First(&wishlist, wishlistId); tx.Error != nil {
		return MakeApiError(404, tx.Error.Error())
	}

	product := &Product{
		Title:      body.Title,
		WishlistId: wishlist.ID,
		Image:      body.Image,
	}

	if tx := Db.Create(product); tx.Error != nil {
		return MakeApiError(500, tx.Error.Error())
	}

	wishlist.Products = append(wishlist.Products, *product)

	return c.Status(201).JSON(wishlist)
}
