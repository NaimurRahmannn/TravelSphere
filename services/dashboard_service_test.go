package services

import (
	"testing"

	"TravelSphere/models"
)

func TestSummarize(t *testing.T) {
	items := []models.WishlistItem{
		{ID: "1", CountryName: "France", Status: models.StatusPlanned},
		{ID: "2", CountryName: "Japan", Status: models.StatusVisited},
		{ID: "3", CountryName: "Brazil", Status: models.StatusPlanned},
		{ID: "4", CountryName: "Egypt", Status: "Unknown"},
	}

	got := summarize(items)

	if got.Total != 4 {
		t.Errorf("Total = %d, want 4", got.Total)
	}
	if got.Planned != 2 {
		t.Errorf("Planned = %d, want 2", got.Planned)
	}
	if got.Visited != 1 {
		t.Errorf("Visited = %d, want 1", got.Visited)
	}
}

func TestSummarize_Empty(t *testing.T) {
	got := summarize([]models.WishlistItem{})
	if got.Total != 0 || got.Planned != 0 || got.Visited != 0 {
		t.Errorf("expected all zero, got %+v", got)
	}
}