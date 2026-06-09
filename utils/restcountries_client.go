package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"TravelSphere/models"
)

const restCountriesBase = "https://restcountries.com/v3.1"
// Shared client for external API calls.
var httpClient = &http.Client{Timeout: 10 * time.Second}


func fetchCountries(url string) ([]models.RawCountry, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("country not found")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body failed: %w", err)
	}

	var raw []models.RawCountry
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, fmt.Errorf("decode failed: %w", err)
	}
	return raw, nil
}

func GetAllCountries() ([]models.RawCountry, error) {
	url := restCountriesBase + "/all?fields=name,capital,population,region,subregion,flags,currencies,languages,latlng,cca2"
	return fetchCountries(url)
}

func GetCountryByName(name string) ([]models.RawCountry, error) {
	url := fmt.Sprintf("%s/name/%s?fields=name,capital,population,region,subregion,flags,currencies,languages,latlng,cca2", restCountriesBase, name)
	return fetchCountries(url)
}