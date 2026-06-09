package controllers

import "TravelSphere/services"

type DashboardController struct {
	BaseController
}

func (c *DashboardController) Get() {
	c.Data["Title"] = "Dashboard"
	c.Data["ActiveNav"] = "dashboard"

	username, _ := c.GetSession("username").(string)

	summary, err := services.GetDashboardSummary(username)
	if err != nil {
		c.Data["LoadError"] = "Could not load dashboard stats."
	} else {
		c.Data["Summary"] = summary
	}

	items, itemErr := services.GetWishlist(username)
	if itemErr == nil {
		c.Data["Items"] = items
	}

	c.Layout = "layout.tpl"
	c.TplName = "dashboard.tpl"
}
