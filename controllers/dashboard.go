package controllers

type DashboardController struct {
	BaseController
}

// GET /wishlist
func (c *DashboardController) Get() {

	c.Layout = "layout.tpl"
	c.TplName = "dashboard.tpl"
}
