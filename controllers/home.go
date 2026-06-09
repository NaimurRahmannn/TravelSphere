package controllers

import (
	"TravelSphere/models"
	"TravelSphere/services"
)

type HomeController struct {
	BaseController
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
