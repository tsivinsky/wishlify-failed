package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tsivinsky/array"
)

func HandleRemoveProduct(c *fiber.Ctx) error {
	wishlistId, err := c.ParamsInt("wishlistId")
	if err != nil {
		return MakeApiError(400, err.Error())
	}

	productId, err := c.ParamsInt("productId")
	if err != nil {
		return MakeApiError(400, err.Error())
	}

	var product *Product
	if tx := Db.First(&product, productId); tx.Error != nil {
		return MakeApiError(404, tx.Error.Error())
	}

	var wishlist *Wishlist
	if tx := Db.Preload("User").Preload("Products").First(&wishlist, wishlistId); tx.Error != nil {
		return MakeApiError(404, tx.Error.Error())
	}

	wishlist.Products = array.Filter(wishlist.Products, func(item Product, i int) bool {
		return item.ID != uint(productId)
	})

	if tx := Db.Delete(&product); tx.Error != nil {
		return MakeApiError(500, tx.Error.Error())
	}

	return c.JSON(wishlist)
}
