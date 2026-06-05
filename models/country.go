package models

// Country is our clean DTO — controllers and templates will use this DTO.
type Country struct {
	Name       string   // common name, e.g. "Bangladesh"
	OfficialName string // e.g. "People's Republic of Bangladesh"
	Slug       string   // url-safe, e.g. "bangladesh"
	Capital    string   // first capital, or "" if none
	Population int      // raw number; format for display separately
	Region     string   // e.g. "Asia"
	Subregion  string   // e.g. "Southern Asia"
	FlagPNG    string   // image URL
	FlagAlt    string   // accessibility text
	Currencies []string // display strings, e.g. "BDT (Bangladeshi taka)"
	Languages  []string // e.g. ["Bengali"]
	LatLng     [2]float64 // [lat, lng] — needed for OpenTripMap attractions
}

// RawCountry mirrors the REST Countries v3.1 JSON shape.
// detects the field that we need
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
	// Currencies is a map keyed by currency code (e.g "BDT"), value(name,symbol)
	Currencies map[string]struct {
		Name   string `json:"name"`
		Symbol string `json:"symbol"`
	} `json:"currencies"`
	// Languages is a map keyed by code (e.g. "ben"), value(language name)
	Languages map[string]string `json:"languages"`
	LatLng    []float64         `json:"latlng"`
	CCA2      string            `json:"cca2"` // 2-letter code, useful fallback for slug
}