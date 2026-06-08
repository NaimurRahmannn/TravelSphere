package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"TravelSphere/models"
	"TravelSphere/utils"
)

// GetWishlist returns every saved entry that belongs to the given user.
// The store keeps all users' items in one file, so we filter by username here.
func GetWishlist(username string) ([]models.WishlistItem, error) {
	all, err := utils.ReadWishlist()
	if err != nil {
		return nil, err
	}

	// Start from an empty (non-nil) slice so the API serialises [] rather than
	// null when the user has no entries yet.
	mine := []models.WishlistItem{}
	for _, item := range all {
		if item.Username == username {
			mine = append(mine, item)
		}
	}
	return mine, nil
}

// CreateWishlistItem validates input, assigns an id and timestamp, tags the
// entry with its owner, and persists.
func CreateWishlistItem(username, countryName, note, status string) (models.WishlistItem, error) {
	if err := utils.ValidateWishlistInput(countryName, status); err != nil {
		return models.WishlistItem{}, err
	}

	items, err := utils.ReadWishlist()
	if err != nil {
		return models.WishlistItem{}, err
	}

	item := models.WishlistItem{
		ID:          uuid.NewString(),
		Username:    username,
		CountryName: countryName,
		Note:        note,
		Status:      status,
		CreatedAt:   time.Now().UTC().Format(time.RFC3339),
	}

	items = append(items, item)
	if err := utils.WriteWishlist(items); err != nil {
		return models.WishlistItem{}, err
	}
	return item, nil
}

// UpdateWishlistItem changes the note and status of an existing entry by id,
// but only if the entry belongs to the given user.
func UpdateWishlistItem(username, id, note, status string) (models.WishlistItem, error) {
	items, err := utils.ReadWishlist()
	if err != nil {
		return models.WishlistItem{}, err
	}

	for i := range items {
		if items[i].ID == id && items[i].Username == username {
			// Validate against the existing country name, since update payloads
			if err := utils.ValidateWishlistInput(items[i].CountryName, status); err != nil {
				return models.WishlistItem{}, err
			}
			items[i].Note = note
			items[i].Status = status
			if err := utils.WriteWishlist(items); err != nil {
				return models.WishlistItem{}, err
			}
			return items[i], nil
		}
	}
	return models.WishlistItem{}, fmt.Errorf("wishlist item %q not found", id)
}

// DeleteWishlistItem removes an entry by id when it belongs to the given user.
// A missing id (or one owned by someone else) is reported as an error so the
// controller can return 404 rather than a silent success.
func DeleteWishlistItem(username, id string) error {
	items, err := utils.ReadWishlist()
	if err != nil {
		return err
	}

	for i := range items {
		if items[i].ID == id && items[i].Username == username {
			items = append(items[:i], items[i+1:]...)
			return utils.WriteWishlist(items)
		}
	}
	return fmt.Errorf("wishlist item %q not found", id)
}

// GetWishlistItemByID returns a single entry by id, scoped to its owner.
func GetWishlistItemByID(username, id string) (models.WishlistItem, error) {
	all, err := utils.ReadWishlist()
	if err != nil {
		return models.WishlistItem{}, err
	}

	for _, item := range all {
		if item.ID == id && item.Username == username {
			return item, nil
		}
	}

	return models.WishlistItem{}, fmt.Errorf("wishlist item not found")
}
