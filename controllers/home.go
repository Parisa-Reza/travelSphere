package controllers

import (
	"travelSphere/models"
	"encoding/json"
	"net/http"
	"time"
	"github.com/beego/beego/v2/server/web"
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
	resp, err := client.Get("https://restcountries.com/v3.1/all?fields=name,capital,flags")
	
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

}