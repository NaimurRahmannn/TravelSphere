package filters

import (
	beecontext "github.com/beego/beego/v2/server/web/context"

	"TravelSphere/utils"
)

//without a session user requireAuth blocks access to protected SSR pages
func RequireAuth(ctx *beecontext.Context) {
	if ctx.Input.Session("username") == nil {
		ctx.Redirect(302, "/login")
	}
}

// RequireAuthAPI guards JSON API
func RequireAuthAPI(ctx *beecontext.Context) {
	if ctx.Input.Session("username") == nil {
		ctx.Output.SetStatus(401)
		_ = ctx.Output.JSON(utils.NewError("authentication required", 401), false, false)
	}
}