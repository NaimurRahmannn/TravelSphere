package filters

import (
	beecontext "github.com/beego/beego/v2/server/web/context"

	"TravelSphere/utils"
)

// RequireAuth blocks access to protected SSR pages for users without a session.
func RequireAuth(ctx *beecontext.Context) {
	// GetSession returns nil when the key isn't set, the user is a guest.
	if ctx.Input.Session("username") == nil {
		// Redirect to login
		ctx.Redirect(302, "/login")
	}
}

// RequireAuthAPI guards the JSON API. Unlike RequireAuth it returns a 401 with a JSON body instead of redirecting: the wishlist endpoints are called via fetch
func RequireAuthAPI(ctx *beecontext.Context) {
	if ctx.Input.Session("username") == nil {
		ctx.Output.SetStatus(401)
		_ = ctx.Output.JSON(utils.NewError("authentication required", 401), false, false)
	}
}