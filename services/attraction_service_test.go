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

// setupAttractionMockServer spins up a fake local server in memory to mimic the OpenTripMap API locally
func setupAttractionMockServer(responseCode int, responseBody interface{}) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(responseCode)
		if responseBody != nil {
			_ = json.NewEncoder(w).Encode(responseBody)
		}
	}))

	// forcing Beego to point its API config variables a into local fake server.
	_ = web.AppConfig.Set("opentripmapkey", "mock_secret_api_key_string")
	_ = web.AppConfig.Set("opentripmapurl", server.URL)
	return server
}

// Testing for GetPopularAttractions with edge cases

func TestGetPopularAttractions_Success(t *testing.T) {
	mockRawList := []models.Attraction{
		{
			Name:  "Eiffel Tower",
			Kinds: "monuments,historical_places,architecture",
		},
		{
			Name:  "  Eiffel Tower  ", // Duplication test
			Kinds: "ignored,tags",
		},
		{
			Name:  "", // Empty check
			Kinds: "nature",
		},
		{
			Name:  "Louvre Museum",
			Kinds: "museums,arts_galleries", //  two kinds check
		},
		{
			Name:  "Notre-Dame",
			Kinds: "churches", // Single kind check
		},
	}

	server := setupAttractionMockServer(http.StatusOK, mockRawList)
	defer server.Close()

	svc := &AttractionService{}

	// Running the main service method to fetch attractions.
	results, err := svc.GetPopularAttractions()
	assert.NoError(t, err)

	// Out of the 5 raw items, only 3 unique, non-empty locations should survive the filter.
	assert.Len(t, results, 3)

	assert.Equal(t, "Eiffel Tower", results[0].Name)
	assert.Equal(t, "monuments, historical places", results[0].DisplayKinds)

	assert.Equal(t, "Louvre Museum", results[1].Name)
	assert.Equal(t, "museums, arts galleries", results[1].DisplayKinds)

	assert.Equal(t, "Notre-Dame", results[2].Name)
	assert.Equal(t, "churches", results[2].DisplayKinds)
}

func TestGetPopularAttractions_MissingConfigKey(t *testing.T) {
	//  misconfigured server environment whithout the API secret key inside the app.conf configuration file.
	_ = web.AppConfig.Set("opentripmapkey", "")
	_ = web.AppConfig.Set("opentripmapurl", "http://any-valid-url")

	svc := &AttractionService{}

	results, err := svc.GetPopularAttractions()

	// throwing an error early and refuse to run without a configuration.
	assert.Nil(t, results)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to retrieve opentripmapkey")
}

func TestGetPopularAttractions_NetworkHttpError(t *testing.T) {
	//breaking the target URL
	_ = web.AppConfig.Set("opentripmapkey", "valid_key")
	_ = web.AppConfig.Set("opentripmapurl", "http://unreachable-address-string-pattern-domain")

	svc := &AttractionService{}

	results, err := svc.GetPopularAttractions()

	// The HTTP client will fail to connect and network error
	assert.Nil(t, results)
	assert.Error(t, err)
}

func TestGetPopularAttractions_BadStatusResponse(t *testing.T) {
	// triggering our fake server with 500 error
	// This tests what happens if the external API service is having a bad day and crashes on us.
	server := setupAttractionMockServer(http.StatusInternalServerError, nil)
	defer server.Close()
	svc := &AttractionService{}
	results, err := svc.GetPopularAttractions()

	//  code catcheing the bad status code, returns no data,
	assert.Nil(t, results)
	assert.Error(t, err)
	assert.EqualError(t, err, "attractions API returned bad status: 500")
}
func TestGetPopularAttractions_JsonDecodeError(t *testing.T) {
	// / triggering fake server with "200 OK" but returns complete garbage plain textinstead of an actual JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("plain unparseable raw string content instead of json payload matrix arrays"))
	}))
	_ = web.AppConfig.Set("opentripmapkey", "valid_key")
	_ = web.AppConfig.Set("opentripmapurl", server.URL)
	defer server.Close()

	svc := &AttractionService{}
	results, err := svc.GetPopularAttractions()

	// JSON decoder will fail to map the garbage text to our structs.
	assert.Nil(t, results)
	assert.Error(t, err)
}
