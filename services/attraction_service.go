package services

import (
	"TravelSphere/utils"
)

const (
	defaultRadius = 10000 // 10km radius around the country centre point
	defaultLimit  = 10    // cap attractions per country so the list stays readable
)

// GetAttractionsByCountry fetches nearby attractions using a country's centre
// coordinates. An API failure returns an empty list rather than an error to the
// caller's UI — a missing attractions section shouldn't take down the whole page.
func GetAttractionsByCountry(lat, lng float64) ([]utils.Attraction, error) {
	attractions, err := utils.GetAttractionsByCoords(lat, lng, defaultRadius, defaultLimit)
	if err != nil {
		return []utils.Attraction{}, err
	}
	return attractions, nil
}

// PopularAttraction is a lightweight showcase entry for the landing page. These
// are static highlights, not live API data, so the home page renders instantly without waiting on OpenTripMap.
type PopularAttraction struct {
	Name  string
	Kinds string
}

// GetPopularAttractions returns the curated landmarks shown on the home page.
func GetPopularAttractions() []PopularAttraction {
	return []PopularAttraction{
		{Name: "Eiffel Tower", Kinds: "architecture, historic"},
		{Name: "Grand Canyon", Kinds: "natural"},
		{Name: "Sydney Opera House", Kinds: "architecture, theatre"},
		{Name: "Colosseum", Kinds: "historic, architecture"},
	}
}