package services

import (
	"fmt"
	"sort"
	"strings"

	"TravelSphere/models"
	"TravelSphere/utils"
)

// slugify converts a country name to a url-safe slug, e.g. "United States" -> "united-states".
func slugify(name string) string {
	s := strings.ToLower(strings.TrimSpace(name))
	s = strings.ReplaceAll(s, " ", "-")
	return s
}

// converts one messy RawCountry into clean Country DTO.
func transform(raw models.RawCountry) models.Country {
	c := models.Country{
		Name:         raw.Name.Common,
		OfficialName: raw.Name.Official,
		Slug:         slugify(raw.Name.Common),
		Population:   raw.Population,
		Region:       raw.Region,
		Subregion:    raw.Subregion,
		FlagPNG:      raw.Flags.PNG,
		FlagAlt:      raw.Flags.Alt,
	}

	// Capital is an array; take the first if present.
	if len(raw.Capital) > 0 {
		c.Capital = raw.Capital[0]
	}

	// LatLng for OpenTripMap later.
	if len(raw.LatLng) == 2 {
		c.LatLng = [2]float64{raw.LatLng[0], raw.LatLng[1]}
	}

	// Flatten the currencies map -> ["BDT (Bangladeshi taka)"].
	for code, cur := range raw.Currencies {
		c.Currencies = append(c.Currencies, fmt.Sprintf("%s (%s)", code, cur.Name))
	}
	sort.Strings(c.Currencies) // stable order so output is deterministic (matters for tests)

	// Flatten the languages map -> ["Bengali", "English"].
	for _, lang := range raw.Languages {
		c.Languages = append(c.Languages, lang)
	}
	sort.Strings(c.Languages)

	return c
}

// GetAllCountries returns all countries as clean DTOs, sorted by name.
func GetAllCountries() ([]models.Country, error) {
	raw, err := utils.GetAllCountries()
	if err != nil {
		return nil, err
	}
	countries := make([]models.Country, 0, len(raw))
	for _, r := range raw {
		countries = append(countries, transform(r))
	}
	sort.Slice(countries, func(i, j int) bool {
		return countries[i].Name < countries[j].Name
	})
	return countries, nil
}

// GetCountryBySlug finds a single country by its slug.
func GetCountryBySlug(slug string) (models.Country, error) {
	// Convert slug back to a search name: "united-states" -> "united states".
	name := strings.ReplaceAll(slug, "-", " ")
	raw, err := utils.GetCountryByName(name)
	if err != nil {
		return models.Country{}, err
	}
	if len(raw) == 0 {
		return models.Country{}, fmt.Errorf("no country for slug %q", slug)
	}
	// Name search can return multiple matches; prefer an exact slug match.
	for _, r := range raw {
		if slugify(r.Name.Common) == slug {
			return transform(r), nil
		}
	}
	return transform(raw[0]), nil // fallback to first result
}