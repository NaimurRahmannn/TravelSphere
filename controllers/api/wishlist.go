package api
import(
	"encoding/json"
	"net/http"
	"github.com/beego/beego/v2/server/web"
	"TravelSphere/services"
	"TravelSphere/utils"
)
// WishlistController exposes wishlist CRUD as JSON at /api/wishlist.
type WishlistController struct{
	web.Controller
}
//createPayload is the expected POST body
type createPayload struct{
	CountryName string `json:"country_name"`
	Note string `json:"note"`
	Status string `json:"status"`
}

//updatePayload is the expected PUT body
type updatePayload struct{
	Note string `json:"note"`
    Status string `json:"status"`
}
//respondJson
func (c * WishlistController) respondJSON(status int, data interface{}){
	c.Ctx.Output.SetStatus(status)
	c.Data["json"]=data
	c.ServeJSON()
}

//respondError
func (c * WishlistController) respondError(message string, status int){
	c.respondJSON(status, utils.NewError(message,status))
}

// Get returns all wishlist entries and according to the ID. GET /api/wishlist
func (c *WishlistController) Get() {
	id := c.Ctx.Input.Param(":id")

	if id != "" {
		item, err := services.GetWishlistItemByID(id)
		if err != nil {
			c.respondError(err.Error(), http.StatusNotFound)
			return
		}

		c.respondJSON(http.StatusOK, item)
		return
	}

	items, err := services.GetWishlist()
	if err != nil {
		c.respondError("Could not load wishlist", http.StatusInternalServerError)
		return
	}

	c.respondJSON(http.StatusOK, items)
}


// Post creates a new entry. POST /api/wishlist
func (c *WishlistController) Post() {
	var payload createPayload
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &payload); err != nil {
		c.respondError("invalid JSON body", http.StatusBadRequest)
		return
	}

	item, err := services.CreateWishlistItem(payload.CountryName, payload.Note, payload.Status)
	if err != nil {
		c.respondError(err.Error(), http.StatusBadRequest)
		return
	}
	c.respondJSON(http.StatusCreated, item) // 201 for successful creation
}


// Put updates note/status of an entry.PUT /api/wishlist/:id
func (c *WishlistController) Put() {
	id := c.Ctx.Input.Param(":id")

	var payload updatePayload
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &payload); err != nil {
		c.respondError("invalid JSON body", http.StatusBadRequest)
		return
	}

	item, err := services.UpdateWishlistItem(id, payload.Note, payload.Status)
	if err != nil {
		c.respondError(err.Error(), http.StatusBadRequest)
		return
	}
	c.respondJSON(http.StatusOK, item)
}


// Delete removes an entry. DELETE /api/wishlist/:id
func (c *WishlistController) Delete() {
	id := c.Ctx.Input.Param(":id")

	if err := services.DeleteWishlistItem(id); err != nil {
		c.respondError(err.Error(), http.StatusNotFound)
		return
	}
	c.respondJSON(http.StatusOK, map[string]string{"message": "deleted"})
}