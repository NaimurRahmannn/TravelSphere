package controllers

import(
	"TravelSphere/services"
)

type CountryController struct{
	BaseController
}

func(c *CountryController)Get(){
	c.Data["Title"]= "Country Explorer"
	c.Data["ActiveNav"]="countries"
	countries,err:=services.GetAllCountries()
	if err!=nil{
	   c.Data["LoadError"]="Could not load countries right now. Please try again shortly."
	   c.Data["Countries"]=[]interface{}{}
	}else{
	   c.Data["Countries"]=countries
	}
	c.Data["Regions"] = []string{"Africa", "Americas", "Asia", "Europe", "Oceania"}

	c.Layout = "layout.tpl"
	c.TplName = "countries.tpl"
}
func (c *CountryController) Detail() {
	slug := c.Ctx.Input.Param(":slug")

	country, err := services.GetCountryBySlug(slug)
	if err != nil {
		// Unknown or malformed slug — show the friendly 404 page, not a stack trace.
		c.renderNotFound(slug)
		return
	}

	c.Data["Title"] = country.Name
	c.Data["ActiveNav"] = "countries"
	c.Data["Country"] = country

	attractions, attrErr := services.GetAttractionsByCountry(country.LatLng[0], country.LatLng[1])
	if attrErr != nil {
		c.Data["AttractionError"] = "Attractions are unavailable right now."
	}
	c.Data["Attractions"] = attractions

	c.Layout = "layout.tpl"
	c.TplName = "destination.tpl"
}

// renderNotFound shows the 404 template with the bad slug echoed back
func (c *CountryController) renderNotFound(slug string) {
	c.Ctx.Output.SetStatus(404)

	c.Data["Title"] = "Not Found"
	c.Data["ActiveNav"] = "countries"
	c.Data["BadSlug"] = slug

	c.Layout = "layout.tpl"
	c.TplName = "404.tpl"

	if err := c.Render(); err != nil {
		c.Ctx.WriteString("Country not found")
	}

	c.StopRun()
}