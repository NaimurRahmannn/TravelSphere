package controllers

import "github.com/beego/beego/v2/server/web"

type BaseController struct {
	web.Controller
}


func (c *BaseController) Prepare() {
	// Shared layout
	c.Layout = "layout.tpl"
	
	// Controllers can override this
	c.Data["ActiveNav"] = ""

	
	isLoggedIn := false
	if user := c.GetSession("username"); user != nil {
		isLoggedIn = true
		c.Data["Username"] = user
	}
	c.Data["IsLoggedIn"] = isLoggedIn

	c.Data["Title"] = "TravelSphere"
}