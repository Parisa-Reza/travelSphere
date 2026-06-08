package api

import (
	"travelSphere/services"

	"github.com/beego/beego/v2/server/web"
)

type CountryAPIController struct {
	web.Controller
}

// for both homepage search suggestions and countries page filtering 

func (c *CountryAPIController) Get() {
	// Check if the "region" query parameter is present to determine the type of request
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
		// Request from homepage
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
