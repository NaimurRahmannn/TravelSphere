package utils

import (
	"os"
	"path/filepath"
	"testing"

	"TravelSphere/models"
)

func redirectStore(t *testing.T) {
	t.Helper()
	original := wishlistFile
	wishlistFile = filepath.Join(t.TempDir(), "wishlist.json")
	t.Cleanup(func() { wishlistFile = original })
}

func TestReadWishlist_MissingFile(t *testing.T) {
	redirectStore(t)

	items, err := ReadWishlist()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(items) != 0 {
		t.Errorf("expected empty list, got %d items", len(items))
	}
}

func TestWriteThenReadWishlist(t *testing.T) {
	redirectStore(t)

	want := []models.WishlistItem{
		{ID: "1", CountryName: "France", Note: "spring", Status: "Planned", CreatedAt: "2026-01-01T00:00:00Z"},
		{ID: "2", CountryName: "Japan", Note: "", Status: "Visited", CreatedAt: "2026-01-02T00:00:00Z"},
	}

	if err := WriteWishlist(want); err != nil {
		t.Fatalf("write failed: %v", err)
	}

	got, err := ReadWishlist()
	if err != nil {
		t.Fatalf("read failed: %v", err)
	}
	if len(got) != 2 {
		t.Fatalf("expected 2 items, got %d", len(got))
	}
	if got[0].CountryName != "France" || got[1].Status != "Visited" {
		t.Errorf("round-trip data mismatch: %+v", got)
	}
}

func TestReadWishlist_EmptyFile(t *testing.T) {
	redirectStore(t)
	if err := os.WriteFile(wishlistFile, []byte{}, 0o644); err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	items, err := ReadWishlist()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(items) != 0 {
		t.Errorf("expected empty list, got %d", len(items))
	}
}

func TestReadWishlist_CorruptJSON(t *testing.T) {
	redirectStore(t)

	if err := os.WriteFile(wishlistFile, []byte("{not valid json"), 0o644); err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	if _, err := ReadWishlist(); err == nil {
		t.Error("expected a decode error, got nil")
	}
}

func TestReadWishlist_ReadError(t *testing.T) {
	redirectStore(t)
	dir := filepath.Join(t.TempDir(), "iamadir")
	if err := os.Mkdir(dir, 0o755); err != nil {
		t.Fatalf("setup failed: %v", err)
	}
	wishlistFile = dir //reading a directory as a file errors

	if _, err := ReadWishlist(); err == nil {
		t.Error("expected a read error when path is a directory, got nil")
	}
}

func TestWriteWishlist_MkdirError(t *testing.T) {
	redirectStore(t)
	blocker := filepath.Join(t.TempDir(), "blocker")
	if err := os.WriteFile(blocker, []byte("x"), 0o644); err != nil {
		t.Fatalf("setup failed: %v", err)
	}
	wishlistFile = filepath.Join(blocker, "wishlist.json")

	if err := WriteWishlist([]models.WishlistItem{}); err == nil {
		t.Error("expected a mkdir error, got nil")
	}
}