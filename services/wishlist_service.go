package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"TravelSphere/models"
	"TravelSphere/utils"
)


func GetWishlist(username string) ([]models.WishlistItem, error) {
	all, err := utils.ReadWishlist()
	if err != nil {
		return nil, err
	}

	//keep api responses as [] for empty wishlists
	mine := []models.WishlistItem{}
	for _, item := range all {
		if item.Username == username {
			mine = append(mine, item)
		}
	}
	return mine, nil
}


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

func UpdateWishlistItem(username, id, note, status string) (models.WishlistItem, error) {
	items, err := utils.ReadWishlist()
	if err != nil {
		return models.WishlistItem{}, err
	}

	for i := range items {
		if items[i].ID == id && items[i].Username == username {
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
