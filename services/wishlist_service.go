package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"TravelSphere/models"
	"TravelSphere/utils"
)

// GetWishlist returns every saved entry.
func GetWishlist() ([]models.WishlistItem, error) {
	return utils.ReadWishlist()
}

// CreateWishlistItem validates input, assigns an id and timestamp, and persists.
func CreateWishlistItem(countryName, note, status string) (models.WishlistItem, error) {
	if err := utils.ValidateWishlistInput(countryName, status); err != nil {
		return models.WishlistItem{}, err
	}

	items, err := utils.ReadWishlist()
	if err != nil {
		return models.WishlistItem{}, err
	}

	item := models.WishlistItem{
		ID:          uuid.NewString(),
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

// UpdateWishlistItem changes the note and status of an existing entry by id.
func UpdateWishlistItem(id, note, status string) (models.WishlistItem, error) {
	items, err := utils.ReadWishlist()
	if err != nil {
		return models.WishlistItem{}, err
	}

	for i := range items {
		if items[i].ID == id {
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

// DeleteWishlistItem removes an entry by id. A missing id is reported as an error so the controller can return 404 rather than a silent success.
func DeleteWishlistItem(id string) error {
	items, err := utils.ReadWishlist()
	if err != nil {
		return err
	}

	for i := range items {
		if items[i].ID == id {
			items = append(items[:i], items[i+1:]...)
			return utils.WriteWishlist(items)
		}
	}
	return fmt.Errorf("wishlist item %q not found", id)
}

func GetWishlistItemByID(id string) (models.WishlistItem, error) {
	items, err := GetWishlist()
	if err != nil {
		return models.WishlistItem{}, err
	}

	for _, item := range items {
		if item.ID == id {
			return item, nil
		}
	}

	return models.WishlistItem{}, fmt.Errorf("wishlist item not found")
}