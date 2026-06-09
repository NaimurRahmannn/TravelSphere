package controllers

import (
	"github.com/beego/beego/v2/server/web"

	"TravelSphere/services"
)


type AuthController struct {
	web.Controller
}


func (c *AuthController) Get() {
	c.Data["Title"] = "Login"
	c.Layout = "layout.tpl"
	c.TplName = "login.tpl"
}


func (c *AuthController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")

	canonical, err := services.AuthenticateUser(username, password)
	if err != nil {
		c.Data["Title"] = "Login"
		c.Data["Error"] = "Invalid username or password."
		c.Data["FormUsername"] = username
		c.Layout = "layout.tpl"
		c.TplName = "login.tpl"
		return
	}

	//Auth filter checks this session key
	c.SetSession("username", canonical)
	c.Redirect("/dashboard", 302)
}


func (c *AuthController) RegisterForm() {
	c.Data["Title"] = "Register"
	c.Layout = "layout.tpl"
	c.TplName = "register.tpl"
}

func (c *AuthController) Register() {
	username := c.GetString("username")
	password := c.GetString("password")
	confirm := c.GetString("confirm")

	if password != confirm {
		c.renderRegisterError(username, "Passwords do not match.")
		return
	}

	user, err := services.RegisterUser(username, password)
	if err != nil {
		c.renderRegisterError(username, err.Error())
		return
	}

	c.SetSession("username", user.Username)
	c.Redirect("/dashboard", 302)
}


func (c *AuthController) renderRegisterError(username, message string) {
	c.Data["Title"] = "Register"
	c.Data["Error"] = message
	c.Data["FormUsername"] = username
	c.Layout = "layout.tpl"
	c.TplName = "register.tpl"
}

func (c *AuthController) Logout() {
	c.DelSession("username")
	c.Redirect("/", 302)
}
