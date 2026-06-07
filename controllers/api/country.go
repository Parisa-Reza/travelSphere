package api

import (
	"travelSphere/services"

	"github.com/beego/beego/v2/server/web"
)

type CountryAPIController struct {
	web.Controller
}

func (c *CountryAPIController) Get() {
	// Check if this request is from the countries page (has region param) or autocomplete (no region param)
	_, hasRegionParam := c.Ctx.Request.URL.Query()["region"]

	if hasRegionParam {
		// Request from countries page - return full country data
		search := c.GetString("search")
		region := c.GetString("region")

		svc := &services.CountryService{}
		countries, err := svc.GetFilteredCountries(search, region)

		if err != nil {
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{"error": "Failed to fetch country data"}
			c.ServeJSON()
			return
		}
		c.Data["json"] = countries
		c.ServeJSON()
	} else {
		// Request from autocomplete - return simplified data
		searchQuery := c.GetString("search")

		svc := &services.CountryService{}
		matches, err := svc.SearchCountries(searchQuery)

		if err != nil {
			c.Ctx.Output.SetStatus(500)
			c.Data["json"] = map[string]string{"error": "Failed to fetch country data"}
			c.ServeJSON()
			return
		}
		c.Data["json"] = matches
		c.ServeJSON()
	}
}
