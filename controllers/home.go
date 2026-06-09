package controllers

import (
	"travelSphere/models"
	"travelSphere/services"
	"encoding/json"
	"net/http"
	"time"
	"github.com/beego/beego/v2/server/web"
	"log"
	"fmt"
)

type HomeController struct {
	web.Controller
}

// for ssr rendering of homepage

func (c *HomeController) Get() {
	// Configure core template layout settings
	c.Layout = "layout.tpl"
	c.TplName = "home.tpl"

	//  Query the RestCountries API dynamically
	
	client := &http.Client{Timeout: 5 * time.Second}

baseUrl, err := web.AppConfig.String("restcountriesurl")

apiUrl := fmt.Sprintf("%s/all?fields=name,capital,flags", baseUrl)

resp, err := client.Get(apiUrl)
	
	var featuredList []models.CountryInfo

	if err == nil && resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()

		var allCountries []models.CountryInfo
		if json.NewDecoder(resp.Body).Decode(&allCountries) == nil {
			// Loop through the API response dynamically and capture the first 8 countries into a slice
			for _, country := range allCountries {
				featuredList = append(featuredList, country)
				
				// Break the loop exactly when we reach 8 entries
				if len(featuredList) == 8 {
					break
				}
			}
		}
	}

	// if the external API fails, bind an empty slice to keep the page rendering safely
	if len(featuredList) == 0 {
		featuredList = []models.CountryInfo{}
	}

	c.Data["FeaturedCountries"] = featuredList


	
	countrySvc := &services.CountryService{}
	featuredCountries, err := countrySvc.GetFilteredCountries("", "")
	if err == nil && len(featuredCountries) > 6 {
		c.Data["FeaturedCountries"] = featuredCountries[:6]
	} else {
		c.Data["FeaturedCountries"] = featuredCountries
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