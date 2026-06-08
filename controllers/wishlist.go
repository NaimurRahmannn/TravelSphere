package controllers

import (
	"TravelSphere/services"
)

type WishlistController struct {
	BaseController
}
// WishlistController render the wishlist page
func (c *WishlistController) Get() {
	c.Data["Title"] = "Wishlist"
	c.Data["ActiveNav"] = "wishlist"
	items, err := services.GetWishlist()
	if err != nil {
		c.Data["LoadError"] = "Could Not load your wishlist right now"
		c.Data["items"] = nil
	} else {
		c.Data["items"] = items
	}
	c.Layout = "layout.tpl"
	c.TplName = "wishlist.tpl"
}
