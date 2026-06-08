package api

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"

	"TravelSphere/services"
	"TravelSphere/utils"
)

//For AJAX refresh DashboardController serves the dashboard counters as JSON
type DashboardController struct {
	web.Controller
}

// Get returns the wishlist summary counts. GET /api/dashboard/summary
func (c *DashboardController) Get() {
	// RequireAuthAPI guards this route, so a username is always present.
	username, _ := c.GetSession("username").(string)
	summary, err := services.GetDashboardSummary(username)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = utils.NewError("could not load summary", http.StatusInternalServerError)
		c.ServeJSON()
		return
	}
	c.Data["json"] = summary
	c.ServeJSON()
}