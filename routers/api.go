package routers

import (
	"TravelSphere/controllers/api"

	"github.com/beego/beego/v2/server/web"
)

// init registers the JSON API routes. Kept separate from the SSR routes in
// router.go so the two route groups stay clearly divided.
func init() {
	// One controller, mapped verb-by-verb. Beego routes GET/POST to the bare
	// path and PUT/DELETE to the :id variant.
	web.Router("/api/wishlist", &api.WishlistController{})
	web.Router("/api/wishlist/:id", &api.WishlistController{})
}