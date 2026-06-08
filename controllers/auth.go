package controllers

import (
	"github.com/beego/beego/v2/server/web"

	"TravelSphere/services"
)


type AuthController struct {
	web.Controller
}

// Get renders the login form GET /login
func (c *AuthController) Get() {
	c.Data["Title"] = "Login"
	c.Layout = "layout.tpl"
	c.TplName = "login.tpl"
}

// Post checks credentials against the user store and starts a session on success
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

	// Session presence is what the auth filter checks for.
	c.SetSession("username", canonical)
	c.Redirect("/dashboard", 302)
}

// RegisterForm renders the registration form-- GET /register
func (c *AuthController) RegisterForm() {
	c.Data["Title"] = "Register"
	c.Layout = "layout.tpl"
	c.TplName = "register.tpl"
}

// Register creates an account and logs the new user straight in---POST /register
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

// renderRegisterError re-renders the registration form with an error message
func (c *AuthController) renderRegisterError(username, message string) {
	c.Data["Title"] = "Register"
	c.Data["Error"] = message
	c.Data["FormUsername"] = username
	c.Layout = "layout.tpl"
	c.TplName = "register.tpl"
}

// Logout clears the session and returns home
func (c *AuthController) Logout() {
	c.DelSession("username")
	c.Redirect("/", 302)
}
