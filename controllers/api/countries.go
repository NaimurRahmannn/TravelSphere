package api

import (
	"net/http"
	"strings"

	"github.com/beego/beego/v2/server/web"

	"TravelSphere/services"
	"TravelSphere/utils"
)

// CountryController serves country data as JSON for the explorer's AJAX  search and region filter.
type CountryController struct {
	web.Controller
}

// Get returns countries, optionally narrowed by ?search= and ?region=. GET /api/countries
func (c *CountryController) Get() {
	search := strings.ToLower(strings.TrimSpace(c.GetString("search")))
	region := strings.TrimSpace(c.GetString("region"))

	all, err := services.GetAllCountries()
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = utils.NewError("could not load countries", http.StatusInternalServerError)
		c.ServeJSON()
		return
	}

	// Filter in the handler for simplicity. For a larger app this belongs in
	// the service, but the matching here is trivial enough to keep inline.
	filtered := all[:0:0] // new zero-length slice, doesn't alias the original
	for _, country := range all {
		// Region filter: skip anything that doesn't match when a region is set.
		if region != "" && country.Region != region {
			continue
		}
		// Search matches name OR capital, case-insensitively.
		if search != "" {
			name := strings.ToLower(country.Name)
			capital := strings.ToLower(country.Capital)
			if !strings.Contains(name, search) && !strings.Contains(capital, search) {
				continue
			}
		}
		filtered = append(filtered, country)
	}

	c.Data["json"] = filtered
	c.ServeJSON()
}