package routers

import (
	"TravelSphere/controllers"
	"TravelSphere/filters"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.InsertFilter("/*", web.BeforeRouter, filters.LogStart)
	web.InsertFilter("/*", web.FinishRouter, filters.LogFinish)
	web.InsertFilter("/wishlist", web.BeforeRouter, filters.RequireAuth)
	web.InsertFilter("/dashboard", web.BeforeRouter, filters.RequireAuth)
	web.Router("/", &controllers.HomeController{})
	web.Router("/countries", &controllers.CountryController{})
	web.Router("/countries/:slug", &controllers.CountryController{}, "get:Detail")
	web.Router("/wishlist", &controllers.WishlistController{})
	web.Router("/dashboard", &controllers.DashboardController{})
	web.Router("/login", &controllers.AuthController{})
	web.Router("/register", &controllers.AuthController{}, "get:RegisterForm;post:Register")
	web.Router("/logout", &controllers.AuthController{}, "get:Logout")
}
