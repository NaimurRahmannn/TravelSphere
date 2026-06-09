package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	beego "github.com/beego/beego/v2/server/web"
)

const samplePlacesJSON = `[
	{"name": "Lalbagh Fort", "kinds": "historic,fortifications", "xid": "Q1", "dist": 1200.5},
	{"name": "", "kinds": "other", "xid": "Q2", "dist": 50},
	{"name": "Ahsan Manzil", "kinds": "museums,historic", "xid": "Q3", "dist": 800}
]`

func setKey(t *testing.T) {
	t.Helper()
	if err := beego.AppConfig.Set("OPENTRIPMAP_API_KEY", "test-key"); err != nil {
		t.Fatalf("could not set test key: %v", err)
	}
}


func withPlacesServer(t *testing.T, status int, body string) {
	t.Helper()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Write([]byte(body))
	}))
	original := openTripMapBase
	openTripMapBase = srv.URL
	t.Cleanup(func() {
		openTripMapBase = original
		srv.Close()
	})
}

func TestGetAttractions_OK(t *testing.T) {
	setKey(t)
	withPlacesServer(t, http.StatusOK, samplePlacesJSON)

	got, err := GetAttractionsByCoords(24, 90, 10000, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// The unnamed place filtered out, leaving 2.
	if len(got) != 2 {
		t.Fatalf("expected 2 named attractions, got %d", len(got))
	}
	if got[0].Name != "Lalbagh Fort" || got[0].Distance != 1200 {
		t.Errorf("unexpected first attraction: %+v", got[0])
	}
}

func TestGetAttractions_ServerError(t *testing.T) {
	setKey(t)
	withPlacesServer(t, http.StatusInternalServerError, `{}`)

	if _, err := GetAttractionsByCoords(24, 90, 10000, 10); err == nil {
		t.Error("expected an error on 500, got nil")
	}
}

func TestGetAttractions_BadJSON(t *testing.T) {
	setKey(t)
	withPlacesServer(t, http.StatusOK, `{not an array`)

	if _, err := GetAttractionsByCoords(24, 90, 10000, 10); err == nil {
		t.Error("expected a decode error, got nil")
	}
}