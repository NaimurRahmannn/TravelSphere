package services

import "TravelSphere/models"

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

func GetDashboardSummary(username string) (DashboardSummary, error) {
	items, err := GetWishlist(username)
	if err != nil {
		return DashboardSummary{}, err
	}
	return summarize(items), nil
}