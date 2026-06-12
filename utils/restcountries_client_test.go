package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

const sampleCountryJSON = `{"data":{"objects":[{
	"names": {"common": "Bangladesh", "official": "People's Republic of Bangladesh"},
	"codes": {"alpha_2": "BD"},
	"capitals": [{"name": "Dhaka"}],
	"population": 169828911,
	"region": "Asia",
	"subregion": "Southern Asia",
	"flag": {"url_png": "https://flagcdn.com/bd.png", "description": "Flag of Bangladesh"},
	"currencies": [{"code": "BDT", "name": "Bangladeshi taka", "symbol": "৳"}],
	"languages": [{"name": "Bengali", "iso639_3": "ben"}],
	"coordinates": {"lat": 24, "lng": 90}
}],"meta":{"more":false}}}`

func withServer(t *testing.T, status int, body string) {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(body))
	}))
	original := restCountriesBase
	restCountriesBase = srv.URL
	t.Cleanup(func() {
		restCountriesBase = original
		srv.Close()
	})
}

func TestGetAllCountries_OK(t *testing.T) {
	withServer(t, http.StatusOK, sampleCountryJSON)

	raw, err := GetAllCountries()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(raw) != 1 {
		t.Fatalf("expected 1 country, got %d", len(raw))
	}
	if raw[0].Name.Common != "Bangladesh" {
		t.Errorf("got name %q, want Bangladesh", raw[0].Name.Common)
	}
}

func TestGetCountryByName_OK(t *testing.T) {
	withServer(t, http.StatusOK, sampleCountryJSON)

	raw, err := GetCountryByName("bangladesh")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(raw) != 1 || raw[0].Name.Common != "Bangladesh" {
		t.Errorf("unexpected result: %+v", raw)
	}
}

func TestFetchCountries_NotFound(t *testing.T) {
	withServer(t, http.StatusNotFound, `{}`)

	if _, err := GetCountryByName("nowhere"); err == nil {
		t.Error("expected a not-found error, got nil")
	}
}

func TestFetchCountries_ServerError(t *testing.T) {
	withServer(t, http.StatusInternalServerError, `{}`)

	if _, err := GetAllCountries(); err == nil {
		t.Error("expected an error on 500, got nil")
	}
}

func TestFetchCountries_BadJSON(t *testing.T) {
	withServer(t, http.StatusOK, `{not an array`)

	if _, err := GetAllCountries(); err == nil {
		t.Error("expected a decode error, got nil")
	}
}