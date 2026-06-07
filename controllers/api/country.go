package api

import (
	"travelSphere/services"

	"github.com/beego/beego/v2/server/web"
)

type CountryAPIController struct {
	web.Controller
}

func (c *CountryAPIController) Get() {
	searchQuery := c.GetString("search")

	service := &services.CountryService{}
	matches, err := service.SearchCountries(searchQuery)

	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to fetch country data"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = matches
	c.ServeJSON()
}
