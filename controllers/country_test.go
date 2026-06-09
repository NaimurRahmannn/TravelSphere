package controllers

import (
	"net/http"
	"testing"
)

func TestCountriesListPage(t *testing.T) {
	w := serve(http.MethodGet, "/countries")
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}
}

func TestCountryDetailFound(t *testing.T) {
	// use a featured slug that exists in the fixtures / external API
	w := serve(http.MethodGet, "/countries/united-states")
	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want 200", w.Code)
	}
}
