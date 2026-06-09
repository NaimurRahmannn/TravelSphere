package utils

import "testing"

func TestValidateWishlistInput(t *testing.T) {
	cases := []struct {
		name        string
		countryName string
		status      string
		wantErr     bool
	}{
		{"valid planned", "Bangladesh", "Planned", false},
		{"valid visited", "France", "Visited", false},
		{"empty country name", "", "Planned", true},
		{"whitespace-only country name", "   ", "Planned", true},
		{"invalid status", "Japan", "Maybe", true},
		{"empty status", "Japan", "", true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := ValidateWishlistInput(c.countryName, c.status)
			if c.wantErr && err == nil {
				t.Errorf("expected an error, got nil")
			}
			if !c.wantErr && err != nil {
				t.Errorf("expected no error, got %v", err)
			}
		})
	}
}