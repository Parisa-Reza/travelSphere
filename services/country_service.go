package services

import (
	"travelSphere/models"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type CountryService struct{}

// SearchCountries fetches and filters countries from the REST API
func (s *CountryService) SearchCountries(searchQuery string) ([]map[string]string, error) {
	searchQuery = strings.ToLower(strings.TrimSpace(searchQuery))
	
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("https://restcountries.com/v3.1/all?fields=name,capital,flags,languages,currency,region")
	
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var dataset []models.CountryInfo
	if err := json.NewDecoder(resp.Body).Decode(&dataset); err != nil {
		return nil, err
	}

	//  Filter the dataset based on the search query and prepare the matches
	var matches []map[string]string
	for _, entry := range dataset {
		if searchQuery != "" && !strings.HasPrefix(strings.ToLower(entry.Name.Common), searchQuery) {
			continue
		}
		
		capitalCity := "N/A"
		if len(entry.Capital) > 0 {
			capitalCity = entry.Capital[0]
		}

		matches = append(matches, map[string]string{
			"label": entry.Name.Common + " — " + capitalCity,
			"slug":  strings.ToLower(entry.Name.Common),
		})
		
		if len(matches) >= 8 {
			break
		}
	}

	return matches, nil
}
