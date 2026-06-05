package controllers

import "github.com/beego/beego/v2/server/web"

// HomeController handles the home page (SSR).
type HomeController struct {
	web.Controller // embedding gives us Get(), Post(), Prepare(), TplName, Data, etc.
}

// Get renders the home page.
func (c *HomeController) Get() {
	c.Data["Title"] = "TravelSphere"
	c.Layout="layout.tpl"
	c.TplName = "home.tpl" 
}