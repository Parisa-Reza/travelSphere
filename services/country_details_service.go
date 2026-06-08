package services

import (
	"fmt"
	"strings"
	"travelSphere/models"
)

type CountryDetailService struct {
	BaseService *CountryService
}

func (s *CountryDetailService) FindBySlug(slug string) (*models.CountryInfo, error) {
	if s.BaseService == nil {
		s.BaseService = &CountryService{}
	}

	list, err := s.BaseService.GetFilteredCountries("", "")
	if err != nil {
		return nil, err
	}

	target := strings.ToLower(strings.TrimSpace(slug))
	for _, item := range list {
		if item.Slug == target {
			return &item, nil
		}
	}

	return nil, fmt.Errorf("destination matching '%s' not found", slug)
}
