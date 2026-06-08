package controllers

import "github.com/beego/beego/v2/server/web"

// AuthController handles login and logout.
type AuthController struct {
	web.Controller
}

// demoUser
const (
	demoUser = "beta"
	demoPass = "beta123"
)

// Get renders the login form. → GET /login
func (c *AuthController) Get() {
	c.Data["Title"] = "Login"
	c.Layout = "layout.tpl"
	c.TplName = "login.tpl"
}

// Post checks credentials and starts a session on success
func (c *AuthController) Post() {
	username := c.GetString("username")
	password := c.GetString("password")

	if username == demoUser && password == demoPass {
		// auth filter checks for session presence.
		c.SetSession("username", username)
		c.Redirect("/dashboard", 302)
		return
	}

	// Re-render the form with an error message; keep it on the same page.
	c.Data["Title"] = "Login"
	c.Data["Error"] = "Invalid username or password."
	c.Layout = "layout.tpl"
	c.TplName = "login.tpl"
}

// Logout clears the session and returns home
func (c *AuthController) Logout() {
	c.DelSession("username")
	c.Redirect("/", 302)
}