package services

import (
	"path/filepath"
	"testing"

	"TravelSphere/utils"
)

func redirectWishlist(t *testing.T) {
	t.Helper()
	restore := utils.SetWishlistFile(filepath.Join(t.TempDir(), "wishlist.json"))
	t.Cleanup(restore)
}

func TestCreateAndGetWishlist(t *testing.T) {
	redirectWishlist(t)

	item, err := CreateWishlistItem("alice", "France", "spring trip", "Planned")
	if err != nil {
		t.Fatalf("create failed: %v", err)
	}
	if item.ID == "" {
		t.Error("expected an ID to be generated")
	}
	if item.Username != "alice" || item.CreatedAt == "" {
		t.Errorf("expected username and timestamp set, got %+v", item)
	}

	mine, err := GetWishlist("alice")
	if err != nil {
		t.Fatalf("get failed: %v", err)
	}
	if len(mine) != 1 || mine[0].CountryName != "France" {
		t.Errorf("expected 1 item for alice, got %v", mine)
	}
}

func TestGetWishlist_OnlyOwnItems(t *testing.T) {
	redirectWishlist(t)

	CreateWishlistItem("alice", "France", "", "Planned")
	CreateWishlistItem("bob", "Japan", "", "Visited")

	aliceItems, _ := GetWishlist("alice")
	if len(aliceItems) != 1 || aliceItems[0].CountryName != "France" {
		t.Errorf("alice should see only her item, got %v", aliceItems)
	}

	bobItems, _ := GetWishlist("bob")
	if len(bobItems) != 1 || bobItems[0].CountryName != "Japan" {
		t.Errorf("bob should see only his item, got %v", bobItems)
	}
}

func TestGetWishlist_EmptyIsNotNil(t *testing.T) {
	redirectWishlist(t)

	items, err := GetWishlist("nobody")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Must be an empty slice, not nil, so the JSON API returns [] not null.
	if items == nil {
		t.Error("expected empty slice, got nil")
	}
	if len(items) != 0 {
		t.Errorf("expected 0 items, got %d", len(items))
	}
}

func TestCreateWishlistItem_ValidationError(t *testing.T) {
	redirectWishlist(t)

	// Empty country name fails validation before any storage write.
	if _, err := CreateWishlistItem("alice", "", "", "Planned"); err == nil {
		t.Error("expected validation error for empty country, got nil")
	}
	// Bad status also fails.
	if _, err := CreateWishlistItem("alice", "France", "", "Maybe"); err == nil {
		t.Error("expected validation error for bad status, got nil")
	}
}

func TestUpdateWishlistItem(t *testing.T) {
	redirectWishlist(t)

	created, _ := CreateWishlistItem("alice", "France", "old note", "Planned")

	updated, err := UpdateWishlistItem("alice", created.ID, "new note", "Visited")
	if err != nil {
		t.Fatalf("update failed: %v", err)
	}
	if updated.Note != "new note" || updated.Status != "Visited" {
		t.Errorf("update didn't apply: %+v", updated)
	}

	// Confirm it persisted.
	got, _ := GetWishlistItemByID("alice", created.ID)
	if got.Note != "new note" {
		t.Errorf("change didn't persist, got note %q", got.Note)
	}
}

func TestUpdateWishlistItem_NotOwner(t *testing.T) {
	redirectWishlist(t)

	created, _ := CreateWishlistItem("alice", "France", "", "Planned")
	if _, err := UpdateWishlistItem("bob", created.ID, "hacked", "Visited"); err == nil {
		t.Error("expected not-found when updating another user's item, got nil")
	}
}

func TestUpdateWishlistItem_BadStatus(t *testing.T) {
	redirectWishlist(t)

	created, _ := CreateWishlistItem("alice", "France", "", "Planned")

	if _, err := UpdateWishlistItem("alice", created.ID, "note", "Nope"); err == nil {
		t.Error("expected validation error for bad status, got nil")
	}
}

func TestDeleteWishlistItem(t *testing.T) {
	redirectWishlist(t)

	created, _ := CreateWishlistItem("alice", "France", "", "Planned")

	if err := DeleteWishlistItem("alice", created.ID); err != nil {
		t.Fatalf("delete failed: %v", err)
	}

	mine, _ := GetWishlist("alice")
	if len(mine) != 0 {
		t.Errorf("expected empty after delete, got %d", len(mine))
	}
}

func TestDeleteWishlistItem_NotFound(t *testing.T) {
	redirectWishlist(t)

	if err := DeleteWishlistItem("alice", "no-such-id"); err == nil {
		t.Error("expected not-found error, got nil")
	}
}

func TestGetWishlistItemByID_NotFound(t *testing.T) {
	redirectWishlist(t)

	if _, err := GetWishlistItemByID("alice", "missing"); err == nil {
		t.Error("expected not-found error, got nil")
	}
}

func TestGetDashboardSummary(t *testing.T) {
	redirectWishlist(t)

	CreateWishlistItem("alice", "France", "", "Planned")
	CreateWishlistItem("alice", "Japan", "", "Visited")
	CreateWishlistItem("alice", "Brazil", "", "Planned")
	CreateWishlistItem("bob", "Egypt", "", "Planned") // bob's, should not count for alice

	summary, err := GetDashboardSummary("alice")
	if err != nil {
		t.Fatalf("summary failed: %v", err)
	}
	if summary.Total != 3 || summary.Planned != 2 || summary.Visited != 1 {
		t.Errorf("unexpected summary for alice: %+v", summary)
	}
}