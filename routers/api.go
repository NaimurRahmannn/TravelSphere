package routers

import (
	"TravelSphere/controllers/api"
	"TravelSphere/filters"

	"github.com/beego/beego/v2/server/web"
)


func init() {
	// The wishlist API is per-user: every entry belongs to a logged-in user, so
	// all wishlist endpoints sit behind the API auth guard (401 JSON for guests).
	web.InsertFilter("/api/wishlist", web.BeforeRouter, filters.RequireAuthAPI)
	web.InsertFilter("/api/wishlist/:id", web.BeforeRouter, filters.RequireAuthAPI)

	// One controller, mapped verb-by-verb. Beego routes GET/POST to the bare path and PUT/DELETE to the :id variant.
	web.Router("/api/wishlist", &api.WishlistController{})
	web.Router("/api/wishlist/:id", &api.WishlistController{})
	web.Router("api/countries",&api.CountryController{})

	// Dashboard counts are derived from the user's own wishlist, so this is guarded too.
	web.InsertFilter("/api/dashboard/summary", web.BeforeRouter, filters.RequireAuthAPI)
	web.Router("/api/dashboard/summary", &api.DashboardController{})
}