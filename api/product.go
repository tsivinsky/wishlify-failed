package main

type Product struct {
	Model

	Title      string   `json:"title" gorm:"title"`
	Image      *string  `json:"image" gorm:"image"`
	WishlistId uint     `json:"wishlistId" gorm:"wishlist_id"`
	Wishlist   Wishlist `json:"wishlist" gorm:"wishlist"`
}
