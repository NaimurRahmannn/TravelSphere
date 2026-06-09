package api

import (
	"net/http"

	"github.com/beego/beego/v2/server/web"

	"TravelSphere/services"
	"TravelSphere/utils"
)


type DashboardController struct {
	web.Controller
}
func (c *DashboardController) Get() {
	// RequireAuthAPI guards this route,a username is always present.
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