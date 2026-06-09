package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"travelSphere/models"
	"github.com/beego/beego/v2/server/web"
)

type CountryService struct{}

// SearchCountries returns country data for autocomplete/search results in homepage search bar
func (s *CountryService) SearchCountries(searchQuery string) ([]map[string]string, error) {
	
searchQuery = strings.ToLower(strings.TrimSpace(searchQuery))

client := &http.Client{Timeout: 5 * time.Second}
	
baseUrl, err := web.AppConfig.String("restcountriesurl")

apiUrl := fmt.Sprintf("%s/all?fields=name,capital", baseUrl)

resp, err := client.Get(apiUrl)
	

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var dataset []models.CountryInfo
	if err := json.NewDecoder(resp.Body).Decode(&dataset); err != nil {
		return nil, err
	}

	var matches []map[string]string

	for _, entry := range dataset {

		capitalCity := "N/A"
		if len(entry.Capital) > 0 {
			capitalCity = entry.Capital[0]
		}

		name := strings.ToLower(entry.Name.Common)
		capital := strings.ToLower(capitalCity)

		// match either country or capital
		if searchQuery != "" &&
			!strings.HasPrefix(name, searchQuery) &&
			!strings.HasPrefix(capital, searchQuery) {
			continue
		}

		matches = append(matches, map[string]string{
			"label": entry.Name.Common + ", " + capitalCity,
			"slug":  name,
		})

		if len(matches) >= 8 {
			break
		}
	}

	if matches == nil {
		matches = []map[string]string{}
	}
	return matches, nil
}

// GetFilteredCountries returns full country data for the countries page with optional filtering
func (s *CountryService) GetFilteredCountries(search, region string) ([]models.CountryInfo, error) {
	client := &http.Client{Timeout: 6 * time.Second}
	

	
baseUrl, err := web.AppConfig.String("restcountriesurl")

apiUrl := fmt.Sprintf("%s/all?fields=name,capital,flags,languages,currencies,region,population", baseUrl)

resp, err := client.Get(apiUrl)


	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var allCountries []models.CountryInfo
	if err := json.NewDecoder(resp.Body).Decode(&allCountries); err != nil {
		return nil, err
	}

	var results []models.CountryInfo
	searchClean := strings.ToLower(strings.TrimSpace(search))
	regionClean := strings.ToLower(strings.TrimSpace(region))

	for i := range allCountries {

		// if capital has multiple entries by taking the first one or if missing marking as "N/A"
		if len(allCountries[i].Capital) > 0 {
			allCountries[i].DisplayCapital = allCountries[i].Capital[0]
		} else {
			allCountries[i].DisplayCapital = "N/A"
		}

		// if languages has multiple entries, we join them with commas for display. If missing, mark as "N/A"
		var langs []string
		for _, lang := range allCountries[i].Languages {
			langs = append(langs, lang)
		}
		allCountries[i].DisplayLanguages = strings.Join(langs, ", ")
		if allCountries[i].DisplayLanguages == "" {
			allCountries[i].DisplayLanguages = "N/A"
		}

		// if currencies has multiple entries, we join them with commas for display. If missing, mark as "N/A"
		var currs []string
		for _, cur := range allCountries[i].Currencies {
			if cur.Name != "" {
				if cur.Symbol != "" {
					currs = append(currs, fmt.Sprintf("%s (%s)", cur.Name, cur.Symbol))
				} else {
					currs = append(currs, cur.Name)
				}
			}
		}
		allCountries[i].DisplayCurrencies = strings.Join(currs, ", ")
		if allCountries[i].DisplayCurrencies == "" {
			allCountries[i].DisplayCurrencies = "N/A"
		}

		// here we create a slug field for each country by lowercasing the name and replacing spaces with hyphens, to be used in URLs
		allCountries[i].Slug = strings.ToLower(strings.ReplaceAll(allCountries[i].Name.Common, " ", "-"))

		// if region filter is applied and the country's region does not match, skip it
		if regionClean != "" && strings.ToLower(allCountries[i].Region) != regionClean {
			continue
		}

		// if search query is applied and it does not match the country's name or capital, we eill skip it
		if searchClean != "" {
			nameMatch := strings.Contains(strings.ToLower(allCountries[i].Name.Common), searchClean)
			capitalMatch := strings.Contains(strings.ToLower(allCountries[i].DisplayCapital), searchClean)
			if !nameMatch && !capitalMatch {
				continue
			}
		}

		results = append(results, allCountries[i])
	}

	if results == nil {
		results = []models.CountryInfo{}
	}
	return results, nil
}
