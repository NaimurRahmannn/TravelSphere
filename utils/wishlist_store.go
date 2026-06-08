package utils

import (
	"TravelSphere/models"
	"encoding/json"
	"os"
	"sync"
)

// wishlistFile is where entries persist between runs. Treated as an external
const wishlistFile = "data/wishlist.json"

// fileMu serialises file access. Beego handles requests concurrently,two simultaneous writes could otherwise clobber each other and corrupt the JSON.
var fileMu sync.Mutex

// ReadWishlist loads all entries from disk.
func ReadWishlist() ([]models.WishlistItem, error) {
	fileMu.Lock()
	defer fileMu.Unlock()
	return readUnlocked()
}

// readUnlocked does the actual read without locking, so write operations that
// already hold the lock can reuse it without deadlocking on themselves.
func readUnlocked() ([]models.WishlistItem, error) {
	bytes, err := os.ReadFile(wishlistFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.WishlistItem{}, nil
		}
		return nil, err
	}
	//empty files also valid
	if len(bytes) == 0 {
		return []models.WishlistItem{}, nil
	}
	var items []models.WishlistItem
	if err := json.Unmarshal(bytes, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func WriteWishlist(items []models.WishlistItem) error {
	fileMu.Lock()
	defer fileMu.Unlock()
	if err := os.MkdirAll("data", 0o755); err != nil {
		return err
	}
	// Indented output keeps the file human-readable if anyone inspects it.
	bytes, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(wishlistFile, bytes, 0o644)
}
