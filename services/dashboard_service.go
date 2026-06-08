package services

import "TravelSphere/models"

// DashboardSummary holds the counts shown on the dashboard. The json tags match what the AJAX refresh expects.
type DashboardSummary struct {
	Total   int `json:"total"`
	Planned int `json:"planned"`
	Visited int `json:"visited"`
}

// summarize counts items by status.
func summarize(items []models.WishlistItem) DashboardSummary {
	summary := DashboardSummary{Total: len(items)}
	for _, item := range items {
		switch item.Status {
		case models.StatusPlanned:
			summary.Planned++
		case models.StatusVisited:
			summary.Visited++
		}
	}
	return summary
}

// reads the wishlist and returns the aggregated counts.
func GetDashboardSummary() (DashboardSummary, error) {
	items, err := GetWishlist()
	if err != nil {
		return DashboardSummary{}, err
	}
	return summarize(items), nil
}