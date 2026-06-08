package models

// Wishlistitems
type WishlistItem struct{
	ID  string `json:"id"`
	Username string `json:"username"`
	CountryName string `json:"country_name"`
	Note  string `json:"note"`
	Status string `json:"status"`
	CreatedAt string `json:"created_at"`
}
//Valid status values. Kept here so the service and validation share the same source
const(
	StatusPlanned="Planned"
	StatusVisited="Visited"
)