package services

import (
	"TravelSphere/utils"
)

const (
	defaultRadius = 10000 // 10 kilometer radius around the country centre point
	defaultLimit  = 10    // cap attractions per country- (the list stays readable)
)

// Attraction failure should not break the country page.
func GetAttractionsByCountry(lat, lng float64) ([]utils.Attraction, error) {
	attractions, err := utils.GetAttractionsByCoords(lat, lng, defaultRadius, defaultLimit)
	if err != nil {
		return []utils.Attraction{}, err
	}
	return attractions, nil
}


type PopularAttraction struct {
	Name  string
	Kinds string
}

// static home page highlights
func GetPopularAttractions() []PopularAttraction {
	return []PopularAttraction{
		{Name: "Eiffel Tower", Kinds: "architecture, historic"},
		{Name: "Grand Canyon", Kinds: "natural"},
		{Name: "Sydney Opera House", Kinds: "architecture, theatre"},
		{Name: "Colosseum", Kinds: "historic, architecture"},
	}
}