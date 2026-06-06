package controllers

import "github.com/beego/beego/v2/server/web"

// BaseController is embedded by all SSR page controllers.
// Its Prepare() runs before every Get()/Post(), setting up shared and template data so individual controllers don't repeat it.

type BaseController struct {
	web.Controller
}

// before the matched HTTP method, Prepare runs automatically
// Use it for request-level setup common to every page.
func (c *BaseController) Prepare() {
	// Shared layout
	c.Layout = "layout.tpl"
	
	// Controllers can override this in their Get() if needed.
	c.Data["ActiveNav"] = ""

	// Login state for the header
	isLoggedIn := false
	if user := c.GetSession("user"); user != nil {
		isLoggedIn = true
		c.Data["Username"] = user
	}
	c.Data["IsLoggedIn"] = isLoggedIn

	// Default title; pages override it.
	c.Data["Title"] = "TravelSphere"
}