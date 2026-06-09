package utils

import "testing"

func TestSetters(t *testing.T) {
	restore := SetRestCountriesBaseURL("http://example.test")
	if restCountriesBase != "http://example.test" {
		t.Error("rest countries base not set")
	}
	restore()

	restore = SetOpenTripMapBaseURL("http://otm.test")
	if openTripMapBase != "http://otm.test" {
		t.Error("opentripmap base not set")
	}
	restore()

	restore = SetWishlistFile("/tmp/x.json")
	if wishlistFile != "/tmp/x.json" {
		t.Error("wishlist file not set")
	}
	restore()

	restore = SetUserFile("/tmp/u.json")
	if userFile != "/tmp/u.json" {
		t.Error("user file not set")
	}
	restore()

	original := httpClient
	restore = SetHTTPClient(nil)
	if httpClient != nil {
		t.Error("http client not set")
	}
	restore()
	if httpClient != original {
		t.Error("http client not restored")
	}
}