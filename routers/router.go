package routers

import (
	"TravelSphere/controllers"

	"github.com/beego/beego/v2/server/web"

)

func init() {
    web.Router("/", &controllers.HomeController{})
	web.Router("/countries", &controllers.CountryController{})
	web.Router("/countries/:slug", &controllers.CountryController{}, "get:Detail")
}
