package controllers

import "github.com/beego/beego/v2/server/web"

type BaseController struct {
	web.Controller
}


func (c *BaseController) Prepare() {
	// Shared layout
	c.Layout = "layout.tpl"
	
	// Controllers can override this in their Get() if needed.
	c.Data["ActiveNav"] = ""

	// Login state for the header. The key must match what AuthController sets on login ("username"), or the header would never show the logged-in state.
	isLoggedIn := false
	if user := c.GetSession("username"); user != nil {
		isLoggedIn = true
		c.Data["Username"] = user
	}
	c.Data["IsLoggedIn"] = isLoggedIn

	// Default title; pages override it.
	c.Data["Title"] = "TravelSphere"
}