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
	// RequireAuth guards this route, so a username is always present here.
	username, _ := c.GetSession("username").(string)
	items, err := services.GetWishlist(username)
	if err != nil {
		c.Data["LoadError"] = "Could Not load your wishlist right now"
		c.Data["Items"] = nil
	} else {
		c.Data["Items"] = items
	}
	c.Layout = "layout.tpl"
	c.TplName = "wishlist.tpl"
}
