package controllers

import (
	"TravelSphere/models"
	"TravelSphere/services"

	"github.com/beego/beego/v2/server/web"
)

// HomeController handles the home page (SSR).
type HomeController struct {
	web.Controller // embedding gives us Get(), Post(), Prepare(), TplName, Data, etc.
}

// featuredSlugs
var featuredSlugs = []string{"united-states", "france", "japan", "australia", "brazil", "bangladesh"}

// Get renders the home page.
func (c *HomeController) Get() {
	c.Data["Title"] = "Home"
	c.Data["ActiveNav"] = "home"
	featured := make([]models.Country, 0, len(featuredSlugs))
	for _, slug := range featuredSlugs {
		country, err := services.GetCountryBySlug(slug)
		if err != nil {
			continue
		}
		featured = append(featured, country)
	}
	c.Data["Featured"] = featured
	c.Data["Featured"] = featured
	c.Layout = "layout.tpl"
	c.TplName = "home.tpl"
}
