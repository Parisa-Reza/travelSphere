package services

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
	"travelSphere/models"
	"travelSphere/utils"

	"github.com/beego/beego/v2/server/web"
)

type CountryService struct{}

// v5 api has been used. deciding whether to fetch data in chunks or all at once.
func (s *CountryService) fetchAllCountries(client *http.Client) ([]models.CountryInfo, error) {
	baseURL, err := web.AppConfig.String("restcountriesurl")
	if err != nil || strings.TrimSpace(baseURL) == "" {
		return nil, fmt.Errorf("restcountriesurl is not configured")
	}

	apiURL := strings.TrimRight(baseURL, "/")
	// If the URL points to V5 and doesn't have a limit already, handle pagination.
	if strings.Contains(apiURL, "api.restcountries.com/countries/v5") && !strings.Contains(apiURL, "limit=") {
		return s.fetchPagedV5Countries(client, apiURL)
	}

	return s.fetchCountriesPage(client, apiURL)
}

// loop through offsets of 100 at a time until the 'more' flag in the JSON tells us we've hit the end.
func (s *CountryService) fetchPagedV5Countries(client *http.Client, apiURL string) ([]models.CountryInfo, error) {
	var allCountries []models.CountryInfo
	for offset := 0; ; offset += 100 {
		pageURL := utils.AddQuery(apiURL, fmt.Sprintf("limit=100&offset=%d", offset))
		countries, more, err := s.fetchV5CountriesPage(client, pageURL)
		if err != nil {
			return nil, err
		}

		allCountries = append(allCountries, countries...)
		if !more { // Break the loop when there are no more pages left.
			break
		}
	}

	if allCountries == nil {
		allCountries = []models.CountryInfo{}
	}
	return allCountries, nil
}

// fetchV5CountriesPage fetches a single block of V5 data and hands the raw body to decoder.
func (s *CountryService) fetchV5CountriesPage(client *http.Client, apiURL string) ([]models.CountryInfo, bool, error) {
	body, err := s.fetchCountryResponse(client, apiURL)
	if err != nil {
		return nil, false, err
	}

	countries, more, isV5, err := utils.DecodeCountriesPage(body)
	if err != nil {
		return nil, false, err
	}
	if !isV5 {
		return countries, false, nil
	}
	return countries, more, nil
}

func (s *CountryService) fetchCountriesPage(client *http.Client, apiURL string) ([]models.CountryInfo, error) {
	body, err := s.fetchCountryResponse(client, apiURL)
	if err != nil {
		return nil, err
	}

	return utils.DecodeCountries(body)
}


// fetchCountryResponse handles building the HTTP GET request, safely injects our Bearer Auth Token from `app.conf` if it exists, and streams down the raw byte response.
func (s *CountryService) fetchCountryResponse(client *http.Client, apiURL string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, err
	}

	apiKey, _ := web.AppConfig.String("restcountrieskey")
	if strings.TrimSpace(apiKey) != "" {
		req.Header.Set("Authorization", "Bearer "+strings.TrimSpace(apiKey))
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (s *CountryService) SearchCountries(searchQuery string) ([]map[string]string, error) {
	searchQuery = strings.ToLower(strings.TrimSpace(searchQuery))

	client := &http.Client{Timeout: 25 * time.Second}
	dataset, err := s.fetchAllCountries(client)
	if err != nil {
		return nil, err
	}

	var matches []map[string]string

	for _, entry := range dataset {
		name := strings.ToLower(entry.Name.Common)
		capital := strings.ToLower(entry.DisplayCapital)

		// We check if the search query matches the start of a country name or capital city.
		if searchQuery != "" &&
			!strings.HasPrefix(name, searchQuery) &&
			!strings.HasPrefix(capital, searchQuery) {
			continue
		}

		matches = append(matches, map[string]string{
			"label": entry.Name.Common + ", " + entry.DisplayCapital,
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

// Used for the main countries page to filter by region or keyword.
func (s *CountryService) GetFilteredCountries(search, region string) ([]models.CountryInfo, error) {
	client := &http.Client{Timeout: 25 * time.Second}
	allCountries, err := s.fetchAllCountries(client)
	if err != nil {
		return nil, err
	}

	var results []models.CountryInfo
	searchClean := strings.ToLower(strings.TrimSpace(search))
	regionClean := strings.ToLower(strings.TrimSpace(region))

	for i := range allCountries {
		// Filter by region if one was selected.
		if regionClean != "" && strings.ToLower(allCountries[i].Region) != regionClean {
			continue
		}

		// Filter by search text if provided.
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
