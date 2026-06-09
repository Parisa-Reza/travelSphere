package controllers

import (
	"travelSphere/services"
	// "github.com/beego/beego/v2/server/web"
)

type CountriesController struct {
	BaseController
}

func (c *CountriesController) Get() {
	// Setup core presentation layout settings
	c.Layout = "layout.tpl"
	c.TplName = "countries.tpl"

	// Get search and region filter from query parameters
	search := c.GetString("search")
	region := c.GetString("region")

	// Fetch filtered countries from service
	service := &services.CountryService{}
	countries, err := service.GetFilteredCountries(search, region)

	if err != nil {
		c.Data["error"] = "Failed to fetch countries"
		return
	}

	c.Data["Countries"] = countries
}
