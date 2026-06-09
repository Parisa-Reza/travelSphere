package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"travelSphere/models"

	"github.com/beego/beego/v2/server/web"
	"github.com/stretchr/testify/assert"
)

// setupDetailsMockServer a fake local API server right in our computer's memory.
func setupDetailsMockServer(responseCode int, responseBody interface{}) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(responseCode)
		if responseBody != nil {
			// Encode mock data to JSON format
			_ = json.NewEncoder(w).Encode(responseBody)
		}
	}))

	// /swapping the real third-party URL in Beego's config with  fake local server URL.
	_ = web.AppConfig.Set("restcountriesurl", server.URL)
	return server
}

// testing for FindBySlug with edge cases

func TestFindBySlug_Success(t *testing.T) {
	//
	// creating some realistic data  expecting back from the API.
	mockData := []models.CountryInfo{
		{
			Name:    models.NameData{Common: "Germany"},
			Capital: []string{"Berlin"},
		},
		{
			Name:    models.NameData{Common: "United Arab Emirates"},
			Capital: []string{"Abu Dhabi"},
		},
	}

	// triggering fake server loaded with the mock data
	server := setupDetailsMockServer(http.StatusOK, mockData)
	defer server.Close()

	svc := &CountryDetailService{}

	// searching for a country using thr slug.
	result, err := svc.FindBySlug("united-arab-emirates")

	// checking the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// checking if the data matches exactly what we loaded into our fake server.
	assert.Equal(t, "United Arab Emirates", result.Name.Common)
	assert.Equal(t, "Abu Dhabi", result.DisplayCapital)

	assert.NotNil(t, svc.BaseService)
}

func TestFindBySlug_NotFoundError(t *testing.T) {
	// putting canada in the fake server
	mockData := []models.CountryInfo{
		{
			Name:    models.NameData{Common: "Canada"},
			Capital: []string{"Ottawa"},
		},
	}

	server := setupDetailsMockServer(http.StatusOK, mockData)
	defer server.Close()

	// /BaseService  ensuring the code handles  safely.
	svc := &CountryDetailService{
		BaseService: &CountryService{},
	}

	// looking for Brazil instead
	result, err := svc.FindBySlug("brazil")

	// "not found" error instead of an empty struct or a total application crash.
	assert.Nil(t, result)
	assert.Error(t, err)
	assert.EqualError(t, err, "destination matching 'brazil' not found")
}

func TestFindBySlug_DependencyServiceError(t *testing.T) {
	// breaking the config URL completely which forces the internal BaseService network request to fail right out of the gate.
	_ = web.AppConfig.Set("restcountriesurl", "http://scrambled-unreachable-network-address-string")

	svc := &CountryDetailService{}

	// Trying to find any slug.
	result, err := svc.FindBySlug("any-slug")

	// Since the service failed, the baseService layerf fail immediately too  and return no data.
	assert.Nil(t, result)
	assert.Error(t, err)
}
