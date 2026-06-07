package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"travelSphere/models"
)

type CountryService struct{}

// SearchCountries returns simplified country data for autocomplete/search results
func (s *CountryService) SearchCountries(searchQuery string) ([]map[string]string, error) {
	searchQuery = strings.ToLower(strings.TrimSpace(searchQuery))

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("https://restcountries.com/v3.1/all?fields=name,capital,flags,languages,currencies,region")

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
		if searchQuery != "" && !strings.HasPrefix(strings.ToLower(entry.Name.Common), searchQuery) {
			continue
		}

		capitalCity := "N/A"
		if len(entry.Capital) > 0 {
			capitalCity = entry.Capital[0]
		}

		matches = append(matches, map[string]string{
			"label": entry.Name.Common + " ," + capitalCity,
			"slug":  strings.ToLower(entry.Name.Common),
		})

		if len(matches) >= 8 {
			break
		}
	}

	// Return empty slice instead of nil to serialize as [] not null
	if matches == nil {
		matches = []map[string]string{}
	}
	return matches, nil
}

// GetFilteredCountries returns full country data for the countries page with optional filtering
func (s *CountryService) GetFilteredCountries(search, region string) ([]models.CountryInfo, error) {
	client := &http.Client{Timeout: 6 * time.Second}
	resp, err := client.Get("https://restcountries.com/v3.1/all?fields=name,capital,flags,languages,currencies,region")

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
		// Flatten Capital Slice safely
		if len(allCountries[i].Capital) > 0 {
			allCountries[i].DisplayCapital = allCountries[i].Capital[0]
		} else {
			allCountries[i].DisplayCapital = "N/A"
		}

		// Flatten Languages map to comma-separated string
		var langs []string
		for _, lang := range allCountries[i].Languages {
			langs = append(langs, lang)
		}
		allCountries[i].DisplayLanguages = strings.Join(langs, ", ")
		if allCountries[i].DisplayLanguages == "" {
			allCountries[i].DisplayLanguages = "N/A"
		}

		// Flatten Currencies map dynamically
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

		// Generate clean navigation lookup slugs
		allCountries[i].Slug = strings.ToLower(strings.ReplaceAll(allCountries[i].Name.Common, " ", "-"))

		// Apply Filters matching query constraints
		if regionClean != "" && strings.ToLower(allCountries[i].Region) != regionClean {
			continue
		}

		if searchClean != "" {
			nameMatch := strings.Contains(strings.ToLower(allCountries[i].Name.Common), searchClean)
			capitalMatch := strings.Contains(strings.ToLower(allCountries[i].DisplayCapital), searchClean)
			if !nameMatch && !capitalMatch {
				continue
			}
		}

		results = append(results, allCountries[i])
	}

	// Return empty slice instead of nil to serialize as [] not null
	if results == nil {
		results = []models.CountryInfo{}
	}
	return results, nil
}
