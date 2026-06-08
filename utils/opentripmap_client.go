package utils

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"io"
	"net/http"
)

const openTripMapBase = "https://api.opentripmap.com/0.1/en/places"

// Attraction is the clean DTO for a single attraction, Used by both the service layer and templates.
type Attraction struct {
	Name     string
	Kinds    string
	XID      string // OpenTripMap unique ID (useful for deduplication)
	Distance int
}

// RawPlace mirrors the OpenTripMap /radius response shape.
type RawPlace struct {
	Name  string  `json:"name"`
	Kinds string  `json:"kinds"`
	XID   string  `json:"xid"`
	Dist  float64 `json:"dist"`
}

// RawPlacesResponse is the top-level response from /radius.
type RawPlacesResponse struct {
	Features []struct {
		Properties RawPlace `json:"properties"`
	} `json:"features"`
}

// GetAttractionsByCoords fetches nearby attractions for a lat/lng.
// radius is in metres. limit caps the number of results.
func GetAttractionsByCoords(lat, lng float64, radius, limit int) ([]Attraction, error) {
	apiKey, err := beego.AppConfig.String("OPENTRIPMAP_API_KEY")
	if err != nil || apiKey == "" {
		return nil, fmt.Errorf("OPENTRIPMAP_API_KEY not set in app.conf")
	}

	url := fmt.Sprintf(
		"%s/radius?radius=%d&lon=%f&lat=%f&limit=%d&format=json&apikey=%s",
		openTripMapBase, radius, lng, lat, limit, apiKey,
	)

	resp, err := httpClient.Get(url) // reuses the same httpClient from restcountries_client.go
	if err != nil {
		return nil, fmt.Errorf("opentripmap request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("opentripmap status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body failed: %w", err)
	}

	// OpenTripMap /radius with format=json returns a plain array, not GeoJSON.
	var raw []RawPlace
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, fmt.Errorf("decode failed: %w", err)
	}

	attractions := make([]Attraction, 0, len(raw))
	for _, p := range raw {
		// Skip unnamed places — they clutter the UI with no value.
		if p.Name == "" {
			continue
		}
		attractions = append(attractions, Attraction{
			Name:     p.Name,
			Kinds:    p.Kinds,
			XID:      p.XID,
			Distance: int(p.Dist),
		})
	}
	return attractions, nil
}
