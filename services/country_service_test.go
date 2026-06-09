package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"TravelSphere/utils"
)

const twoCountriesJSON = `[
	{
		"name": {"common": "Zimbabwe", "official": "Republic of Zimbabwe"},
		"capital": ["Harare"],
		"population": 15000000,
		"region": "Africa",
		"flags": {"png": "z.png", "alt": "flag"},
		"currencies": {"ZWL": {"name": "Zimbabwean dollar"}},
		"languages": {"eng": "English", "sna": "Shona"},
		"latlng": [-20, 30],
		"cca2": "ZW"
	},
	{
		"name": {"common": "Albania", "official": "Republic of Albania"},
		"capital": ["Tirana"],
		"population": 2400000,
		"region": "Europe",
		"flags": {"png": "a.png", "alt": "flag"},
		"currencies": {"ALL": {"name": "Albanian lek"}},
		"languages": {"sqi": "Albanian"},
		"latlng": [41, 20],
		"cca2": "AL"
	}
]`


func startCountryServer(t *testing.T, status int, body string) {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(body))
	}))
	restore := utils.SetRestCountriesBaseURL(srv.URL)
	t.Cleanup(func() {
		restore()
		srv.Close()
	})
}

func TestGetAllCountries_TransformsAndSorts(t *testing.T) {
	startCountryServer(t, http.StatusOK, twoCountriesJSON)

	countries, err := GetAllCountries()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(countries) != 2 {
		t.Fatalf("expected 2 countries, got %d", len(countries))
	}
	if countries[0].Name != "Albania" {
		t.Errorf("expected Albania first after sort, got %q", countries[0].Name)
	}
	if countries[0].Slug != "albania" {
		t.Errorf("expected slug 'albania', got %q", countries[0].Slug)
	}
	if len(countries[0].Currencies) != 1 || countries[0].Currencies[0] != "ALL (Albanian lek)" {
		t.Errorf("unexpected currency formatting: %v", countries[0].Currencies)
	}
	zim := countries[1]
	if len(zim.Languages) != 2 || zim.Languages[0] != "English" || zim.Languages[1] != "Shona" {
		t.Errorf("expected sorted [English Shona], got %v", zim.Languages)
	}
	if countries[0].Capital != "Tirana" {
		t.Errorf("expected capital Tirana, got %q", countries[0].Capital)
	}
	if zim.LatLng != [2]float64{-20, 30} {
		t.Errorf("expected latlng [-20 30], got %v", zim.LatLng)
	}
}

func TestGetAllCountries_ClientError(t *testing.T) {
	startCountryServer(t, http.StatusInternalServerError, `{}`)

	if _, err := GetAllCountries(); err == nil {
		t.Error("expected an error when the client fails, got nil")
	}
}

func TestGetCountryBySlug_ExactMatch(t *testing.T) {
	startCountryServer(t, http.StatusOK, twoCountriesJSON)
	c, err := GetCountryBySlug("albania")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.Name != "Albania" {
		t.Errorf("expected Albania, got %q", c.Name)
	}
}

func TestGetCountryBySlug_FallbackToFirst(t *testing.T) {
	startCountryServer(t, http.StatusOK, twoCountriesJSON)
	c, err := GetCountryBySlug("no-such-slug")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.Name == "" {
		t.Error("expected a fallback country, got empty")
	}
}

func TestGetCountryBySlug_Empty(t *testing.T) {
	startCountryServer(t, http.StatusOK, `[]`)

	if _, err := GetCountryBySlug("anything"); err == nil {
		t.Error("expected an error for an empty result, got nil")
	}
}

func TestGetCountryBySlug_ClientError(t *testing.T) {
	startCountryServer(t, http.StatusInternalServerError, `{}`)

	if _, err := GetCountryBySlug("albania"); err == nil {
		t.Error("expected an error when the client fails, got nil")
	}
}