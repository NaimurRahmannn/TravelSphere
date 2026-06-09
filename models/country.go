package models

type Country struct {
	Name       string   // "Bangladesh"
	OfficialName string //"People's Republic of Bangladesh"
	Slug       string   //  "bangladesh"
	Capital    string   // first capital, or "" if none
	Population int      // raw number; format for display separately
	Region     string   //  "Asia"
	Subregion  string   //"Southern Asia"
	FlagPNG    string   // image_URL
	FlagAlt    string   // flag alternative text
	Currencies []string //  "BDT -Bangladeshi taka"
	Languages  []string //"Bengali"
	LatLng     [2]float64 // lat, lng
}

// RawCountry matches the fields used from the REST Countries API response.
type RawCountry struct {
	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name"`
	Capital    []string `json:"capital"`
	Population int      `json:"population"`
	Region     string   `json:"region"`
	Subregion  string   `json:"subregion"`
	Flags      struct {
		PNG string `json:"png"`
		Alt string `json:"alt"`
	} `json:"flags"`
	Currencies map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	Languages map[string]string `json:"languages"`
	LatLng    []float64         `json:"latlng"`
	CCA2      string            `json:"cca2"` // two-letter code, useful fallback for slug
}