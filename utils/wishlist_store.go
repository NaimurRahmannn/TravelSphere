package utils

import (
	"TravelSphere/models"
	"encoding/json"
	"os"
	"sync"
)

// wishlistFile is where entries persist
const wishlistFile = "data/wishlist.json"


var fileMu sync.Mutex

func ReadWishlist() ([]models.WishlistItem, error) {
	fileMu.Lock()
	defer fileMu.Unlock()
	return readUnlocked()
}

// Used when the caller already holds fileMu.
func readUnlocked() ([]models.WishlistItem, error) {
	bytes, err := os.ReadFile(wishlistFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []models.WishlistItem{}, nil
		}
		return nil, err
	}
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
	bytes, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(wishlistFile, bytes, 0o644)
}
