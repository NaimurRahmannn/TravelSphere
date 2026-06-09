package routers

import (
	"TravelSphere/controllers/api"
	"TravelSphere/filters"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	//Wishlist api requires login.
	web.InsertFilter("/api/wishlist", web.BeforeRouter, filters.RequireAuthAPI)
	web.InsertFilter("/api/wishlist/:id", web.BeforeRouter, filters.RequireAuthAPI)

	web.Router("/api/wishlist", &api.WishlistController{})
	web.Router("/api/wishlist/:id", &api.WishlistController{})
	web.Router("/api/countries", &api.CountryController{})

	// Dashboard summary uses user specific data.
	web.InsertFilter("/api/dashboard/summary", web.BeforeRouter, filters.RequireAuthAPI)
	web.Router("/api/dashboard/summary", &api.DashboardController{})
}
