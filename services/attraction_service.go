package services

// PopularAttraction is a lightweight showcase entry for the landing page.
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