package controllers

import "TravelSphere/services"

// DashboardController renders the dashboard page (SSR).
type DashboardController struct {
	BaseController
}

func (c *DashboardController) Get() {
	c.Data["Title"] = "Dashboard"
	c.Data["ActiveNav"] = "dashboard"

	// RequireAuth guards this route, so a username is always present here.
	username, _ := c.GetSession("username").(string)

	summary, err := services.GetDashboardSummary(username)
	if err != nil {
		c.Data["LoadError"] = "Could not load dashboard stats."
	} else {
		c.Data["Summary"] = summary
	}

	// The saved-destinations list reuses the same wishlist data.
	items, itemErr := services.GetWishlist(username)
	if itemErr == nil {
		c.Data["Items"] = items
	}

	c.Layout = "layout.tpl"
	c.TplName = "dashboard.tpl"
}