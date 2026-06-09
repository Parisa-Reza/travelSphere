package controllers

import (
	"encoding/json"
	"log"
	"travelSphere/services"
)

type WishlistController struct {
	BaseController
}

// GET /wishlist
func (c *WishlistController) Get() {
	userID := c.GetSession("user_id").(string)
	svc := &services.WishlistService{}

	c.Data["WishlistItems"] = svc.GetUserWishlist(userID)
	c.Layout = "layout.tpl"
	c.TplName = "wishlist.tpl"
}

// GET /api/wishlist
func (c *WishlistController) GetAPI() {
	userID := c.GetSession("user_id").(string)
	svc := &services.WishlistService{}
	c.Data["json"] = svc.GetUserWishlist(userID)
	c.ServeJSON()
}

// POST /api/wishlist
func (c *WishlistController) PostAPI() {
	userID := c.GetSession("user_id").(string)

	var body struct {
		CountryName string `json:"name"`
	}

	// Used json.Unmarshal since c.Ctx.Input.RequestBody is a raw []byte slice
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.ServeJSON()
		return
	}

	log.Printf("[DEBUG] Wishlist Add Request - User: %s, Country received: '%s'", userID, body.CountryName)

	if body.CountryName == "" {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Country name parameter was empty or failed to parse"}
		c.ServeJSON()
		return
	}

	svc := &services.WishlistService{}
	entry, err := svc.AddToWishlist(userID, body.CountryName)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = entry
	c.ServeJSON()
}

// PUT /api/wishlist/:id
func (c *WishlistController) PutAPI() {
	userID := c.GetSession("user_id").(string)
	id := c.Ctx.Input.Param(":id")

	var body struct {
		Note   string `json:"note"`
		Status string `json:"status"`
	}

	// Used json.Unmarshal since c.Ctx.Input.RequestBody is a raw []byte slice
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.ServeJSON()
		return
	}

	svc := &services.WishlistService{}
	entry, err := svc.UpdateWishlist(id, userID, body.Note, body.Status)
	if err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = entry
	c.ServeJSON()
}

// DELETE /api/wishlist/:id
func (c *WishlistController) DeleteAPI() {
	userID := c.GetSession("user_id").(string)
	id := c.Ctx.Input.Param(":id")

	svc := &services.WishlistService{}
	if err := svc.DeleteWishlist(id, userID); err != nil {
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]string{"success": "Item deleted"}
	c.ServeJSON()
}
