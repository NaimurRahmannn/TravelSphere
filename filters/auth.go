package filters

import (
	beecontext "github.com/beego/beego/v2/server/web/context"

	"TravelSphere/utils"
)

// without a session, RequireAuth blocks access to protected SSR pages
func RequireAuth(ctx *beecontext.Context) {
	if loggedIn(ctx) {
		return
	}
	ctx.Redirect(302, "/login")
}

// RequireAuthAPI guards JSON API routes
func RequireAuthAPI(ctx *beecontext.Context) {
	if loggedIn(ctx) {
		return
	}
	ctx.Output.SetStatus(401)
	_ = ctx.Output.JSON(utils.NewError("authentication required", 401), false, false)
}
func loggedIn(ctx *beecontext.Context) bool {
	if ctx.Input.CruSession == nil {
		return false
	}
	return ctx.Input.Session("username") != nil
}
