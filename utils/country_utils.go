package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"travelSphere/models"
)

// AddQuery adds query params whether the URL already has a query string or not.
func AddQuery(apiURL, query string) string {
	if strings.Contains(apiURL, "?") {
		return apiURL + "&" + query
	}
	return apiURL + "?" + query
}

func DecodeCountries(body []byte) ([]models.CountryInfo, error) {
	countries, _, _, err := DecodeCountriesPage(body)
	return countries, err
}

// here DecodeCountriesPage maps the nested V5 JSON into the simpler internal CountryInfo format.
func DecodeCountriesPage(body []byte) ([]models.CountryInfo, bool, bool, error) {
	var v5 models.RestCountriesV5Response
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&v5); err == nil && v5.Data.Objects != nil {
		countries := make([]models.CountryInfo, 0, len(v5.Data.Objects))
		for _, entry := range v5.Data.Objects {
			country := models.CountryInfo{
				Name: models.NameData{
					Common:   entry.Names.Common,
					Official: entry.Names.Official,
				},
				Region:     entry.Region,
				Population: entry.Population,
				Languages:  map[string]string{},
				Currencies: map[string]models.CurrencyDetail{},
				Flags: models.FlagData{
					Png: entry.Flag.Png,
					Svg: entry.Flag.Svg,
				},
			}

			for _, capital := range entry.Capitals {
				if capital.Name != "" {
					country.Capital = append(country.Capital, capital.Name)
				}
			}

			for _, lang := range entry.Languages {
				if lang.Name == "" {
					continue
				}
				code := lang.Code
				if code == "" {
					code = lang.Name
				}
				country.Languages[code] = lang.Name
			}

			for _, currency := range entry.Currencies {
				if currency.Name == "" {
					continue
				}
				code := currency.Code
				if code == "" {
					code = currency.Name
				}
				country.Currencies[code] = models.CurrencyDetail{
					Name:   currency.Name,
					Symbol: currency.Symbol,
				}
			}

			countries = append(countries, NormalizeCountry(country))
		}
		return countries, v5.Data.Meta.More, true, nil
	}

	var countries []models.CountryInfo
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&countries); err != nil {
		return nil, false, false, err
	}

	for i := range countries {
		countries[i] = NormalizeCountry(countries[i])
	}
	return countries, false, false, nil
}

// NormalizeCountry fills display-only fields used by templates and JSON responses.
func NormalizeCountry(country models.CountryInfo) models.CountryInfo {
	if len(country.Capital) > 0 {
		country.DisplayCapital = country.Capital[0]
	} else {
		country.DisplayCapital = "N/A"
	}

	var langs []string
	for _, lang := range country.Languages {
		langs = append(langs, lang)
	}
	country.DisplayLanguages = strings.Join(langs, ", ")
	if country.DisplayLanguages == "" {
		country.DisplayLanguages = "N/A"
	}

	var currs []string
	for _, cur := range country.Currencies {
		if cur.Name != "" {
			if cur.Symbol != "" {
				currs = append(currs, fmt.Sprintf("%s (%s)", cur.Name, cur.Symbol))
			} else {
				currs = append(currs, cur.Name)
			}
		}
	}
	country.DisplayCurrencies = strings.Join(currs, ", ")
	if country.DisplayCurrencies == "" {
		country.DisplayCurrencies = "N/A"
	}

	country.Slug = strings.ToLower(strings.ReplaceAll(country.Name.Common, " ", "-"))
	return country
}
