package controllers

import (
	"log"
	"travelSphere/services"
)

type HomeController struct {
	BaseController
}

// for ssr rendering of homepage

func (c *HomeController) Get() {
	// Configure core template layout settings
	c.Layout = "layout.tpl"
	c.TplName = "home.tpl"

	countrySvc := &services.CountryService{}
	featuredCountries, err := countrySvc.GetFilteredCountries("", "")
	if err == nil && len(featuredCountries) > 8 {
		c.Data["FeaturedCountries"] = featuredCountries[:8]
	} else if err == nil {
		c.Data["FeaturedCountries"] = featuredCountries
	} else {
		log.Printf("[ Failed to retrieve external country data: %v", err)
		c.Data["FeaturedCountries"] = nil
	}

	attractionSvc := &services.AttractionService{}
	popularAttractions, err := attractionSvc.GetPopularAttractions()
	if err != nil {
		log.Printf("[Home Error] Failed to retrieve external attractions data: %v", err)
		c.Data["PopularAttractions"] = nil
	} else {
		c.Data["PopularAttractions"] = popularAttractions
	}

}
