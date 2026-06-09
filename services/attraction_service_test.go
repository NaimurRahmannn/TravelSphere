package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	beego "github.com/beego/beego/v2/server/web"

	"TravelSphere/utils"
)

const placesJSON = `[
	{"name": "Lalbagh Fort", "kinds": "historic", "xid": "Q1", "dist": 1000},
	{"name": "Ahsan Manzil", "kinds": "museums", "xid": "Q2", "dist": 500}
]`

func TestGetAttractionsByCountry_OK(t *testing.T) {
	_ = beego.AppConfig.Set("OPENTRIPMAP_API_KEY", "test-key")

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(placesJSON))
	}))
	restore := utils.SetOpenTripMapBaseURL(srv.URL)
	defer restore()
	defer srv.Close()

	got, err := GetAttractionsByCountry(24, 90)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("expected 2 attractions, got %d", len(got))
	}
}

func TestGetAttractionsByCountry_Error(t *testing.T) {
	_ = beego.AppConfig.Set("OPENTRIPMAP_API_KEY", "test-key")

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	restore := utils.SetOpenTripMapBaseURL(srv.URL)
	defer restore()
	defer srv.Close()
	got, err := GetAttractionsByCountry(24, 90)
	if err == nil {
		t.Error("expected an error, got nil")
	}
	if got == nil {
		t.Error("expected a non-nil empty slice on error, got nil")
	}
}

func TestGetPopularAttractions(t *testing.T) {
	got := GetPopularAttractions()
	if len(got) != 4 {
		t.Errorf("expected 4 popular attractions, got %d", len(got))
	}
	if got[0].Name != "Eiffel Tower" {
		t.Errorf("expected Eiffel Tower first, got %q", got[0].Name)
	}
}