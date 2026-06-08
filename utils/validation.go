package utils

import (
	"TravelSphere/models"
	"fmt"
	"strings"
)
// ValidateWishlistInput checks a create/update payload. 
func ValidateWishlistInput(countryName, status string) error {
	if strings.TrimSpace(countryName) == "" {
		return fmt.Errorf("country_name is required")
	}
	if status != models.StatusPlanned && status != models.StatusVisited {
		return fmt.Errorf("status must be %q or %q", models.StatusPlanned, models.StatusVisited)
	}
	return nil
}
