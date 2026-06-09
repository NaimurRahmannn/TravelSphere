package api
import(
	"encoding/json"
	"net/http"
	"github.com/beego/beego/v2/server/web"
	"TravelSphere/services"
	"TravelSphere/utils"
)
//JSON endpoints for wishlist actions.
type WishlistController struct{
	web.Controller
}

type createPayload struct{
	CountryName string `json:"country_name"`
	Note string `json:"note"`
	Status string `json:"status"`
}


type updatePayload struct{
	Note string `json:"note"`
    Status string `json:"status"`
}

func (c * WishlistController) respondJSON(status int, data interface{}){
	c.Ctx.Output.SetStatus(status)
	c.Data["json"]=data
	c.ServeJSON()
}


func (c * WishlistController) respondError(message string, status int){
	c.respondJSON(status, utils.NewError(message,status))
}

//currentUser returns the logged-in username from the session
func (c *WishlistController) currentUser() string {
	username, _ := c.GetSession("username").(string)
	return username
}

func (c *WishlistController) Get() {
	id := c.Ctx.Input.Param(":id")

	if id != "" {
		item, err := services.GetWishlistItemByID(c.currentUser(), id)
		if err != nil {
			c.respondError(err.Error(), http.StatusNotFound)
			return
		}

		c.respondJSON(http.StatusOK, item)
		return
	}

	items, err := services.GetWishlist(c.currentUser())
	if err != nil {
		c.respondError("Could not load wishlist", http.StatusInternalServerError)
		return
	}

	c.respondJSON(http.StatusOK, items)
}

func (c *WishlistController) Post() {
	var payload createPayload
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &payload); err != nil {
		c.respondError("invalid JSON body", http.StatusBadRequest)
		return
	}

	item, err := services.CreateWishlistItem(c.currentUser(), payload.CountryName, payload.Note, payload.Status)
	if err != nil {
		c.respondError(err.Error(), http.StatusBadRequest)
		return
	}
	c.respondJSON(http.StatusCreated, item) 
}

func (c *WishlistController) Put() {
	id := c.Ctx.Input.Param(":id")

	var payload updatePayload
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &payload); err != nil {
		c.respondError("invalid JSON body", http.StatusBadRequest)
		return
	}

	item, err := services.UpdateWishlistItem(c.currentUser(), id, payload.Note, payload.Status)
	if err != nil {
		c.respondError(err.Error(), http.StatusBadRequest)
		return
	}
	c.respondJSON(http.StatusOK, item)
}

func (c *WishlistController) Delete() {
	id := c.Ctx.Input.Param(":id")

	if err := services.DeleteWishlistItem(c.currentUser(), id); err != nil {
		c.respondError(err.Error(), http.StatusNotFound)
		return
	}
	c.respondJSON(http.StatusOK, map[string]string{"message": "deleted"})
}