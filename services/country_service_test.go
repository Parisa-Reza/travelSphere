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

// setupMockServer cretae a local temporary HTTP server containing custom mock response data
func setupMockServer(responseCode int, responseBody interface{}) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(responseCode)
		if responseBody != nil {
			_ = json.NewEncoder(w).Encode(responseBody)
		}
	}))

	_ = web.AppConfig.Set("restcountriesurl", server.URL)
	return server
}

// Tests for SearchCountries

func TestSearchCountries_Success(t *testing.T) {
	// Sample backend datasets mimicking REST Countries return format
	mockData := []models.CountryInfo{
		{
			Name:    models.NameData{Common: "Bangladesh"},
			Capital: []string{"Dhaka"},
		},
		{
			Name:    models.NameData{Common: "Germany"},
			Capital: []string{"Berlin"},
		},
		{
			Name:    models.NameData{Common: "Canada"},
			Capital: []string{}, // No capital entry test case
		},
	}

	server := setupMockServer(http.StatusOK, mockData)
	defer server.Close()

	svc := &CountryService{}

	//  Search with empty query string (should return all matches up to limit)
	results, err := svc.SearchCountries("")
	assert.NoError(t, err)
	assert.Len(t, results, 3)
	assert.Equal(t, "Bangladesh, Dhaka", results[0]["label"])
	assert.Equal(t, "bangladesh", results[0]["slug"])
	assert.Equal(t, "Canada, N/A", results[2]["label"]) // Capital fallback check

	//  Search query matching country
	results, err = svc.SearchCountries("bang")
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, "Bangladesh, Dhaka", results[0]["label"])

	//  Search query matching capital city
	results, err = svc.SearchCountries("berl")
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, "Germany, Berlin", results[0]["label"])

	//  Search query that has no matching element
	results, err = svc.SearchCountries("xyz")
	assert.NoError(t, err)
	assert.Len(t, results, 0)
	assert.NotNil(t, results) // this ensure it returned empty slice rather than nil
}

func TestSearchCountries_LimitEnforcement(t *testing.T) {
	// in the code we limited max suggestion to 8 in the search box

	var largeMockData []models.CountryInfo
	for i := 1; i <= 10; i++ {
		largeMockData = append(largeMockData, models.CountryInfo{
			Name:    models.NameData{Common: "Country"},
			Capital: []string{"Capital"},
		})
	}

	server := setupMockServer(http.StatusOK, largeMockData)
	defer server.Close()

	svc := &CountryService{}
	results, err := svc.SearchCountries("count")
	assert.NoError(t, err)
	assert.Len(t, results, 8) //  execution breaks loop at exactly 8 items
}

// forcing the code to jump straight into its low-level network error-handling path.
func TestSearchCountries_NetworkHttpError(t *testing.T) {

	_ = web.AppConfig.Set("restcountriesurl", "http://invalid-url-string-structure")

	svc := &CountryService{}
	results, err := svc.SearchCountries("test")
	assert.Error(t, err)
	assert.Nil(t, results)
}

// Sending  back  plain text instead of JSON to force a data parsing error.

func TestSearchCountries_JsonParseError(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("invalid json data parsing structural string"))
	}))
	_ = web.AppConfig.Set("restcountriesurl", server.URL)
	defer server.Close()

	svc := &CountryService{}
	results, err := svc.SearchCountries("test")
	assert.Error(t, err)
	assert.Nil(t, results)
}

// Tests for GetFilteredCountries

func TestGetFilteredCountries_Success(t *testing.T) {
	mockData := []models.CountryInfo{
		{
			Name:       models.NameData{Common: "United States"},
			Capital:    []string{"Washington D.C."},
			Region:     "Americas",
			Population: 331000000,
			Languages:  map[string]string{"eng": "English"},
			Currencies: map[string]models.CurrencyDetail{
				"USD": {Name: "United States Dollar", Symbol: "$"},
			},
		},
		{
			Name:       models.NameData{Common: "United Kingdom"},
			Capital:    []string{"London"},
			Region:     "Europe",
			Population: 67000000,
			Languages:  map[string]string{"eng": "English"},
			Currencies: map[string]models.CurrencyDetail{
				"GBP": {Name: "British Pound", Symbol: ""}, // Currency missing symbol verification
			},
		},
		{
			Name:       models.NameData{Common: "EmptyLand"},
			Capital:    []string{},
			Region:     "Asia",
			Population: 100,
			Languages:  map[string]string{},
			Currencies: map[string]models.CurrencyDetail{
				"BLANK": {Name: "", Symbol: ""}, // Currency missing name verification
			},
		},
	}

	server := setupMockServer(http.StatusOK, mockData)
	defer server.Close()

	svc := &CountryService{}

	// fetching without passing any region or input in the search box
	results, err := svc.GetFilteredCountries("", "")
	assert.NoError(t, err)
	assert.Len(t, results, 3)

	// Check transformations and formatting mutations
	assert.Equal(t, "Washington D.C.", results[0].DisplayCapital)
	assert.Equal(t, "English", results[0].DisplayLanguages)
	assert.Equal(t, "United States Dollar ($)", results[0].DisplayCurrencies)
	assert.Equal(t, "united-states", results[0].Slug)

	// if any structure missing
	assert.Equal(t, "N/A", results[2].DisplayCapital)
	assert.Equal(t, "N/A", results[2].DisplayLanguages)
	assert.Equal(t, "N/A", results[2].DisplayCurrencies)

	// filter countries by regoin
	results, err = svc.GetFilteredCountries("", "europe")
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, "United Kingdom", results[0].Name.Common)

	// filter countries by search query match on country
	results, err = svc.GetFilteredCountries("states", "")
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, "United States", results[0].Name.Common)

	// filter countries by search query match on capital
	results, err = svc.GetFilteredCountries("london", "")
	assert.NoError(t, err)
	assert.Len(t, results, 1)
	assert.Equal(t, "United Kingdom", results[0].Name.Common)

	//  searching  term doesn't match properties
	results, err = svc.GetFilteredCountries("nonexistent-search-term", "")
	assert.NoError(t, err)
	assert.Len(t, results, 0)
	assert.NotNil(t, results)
}

//  //We want to test what happens if the internet completely drops.To mock a total network collapse, we mess up the configuration URL on purpose.By pointing it to a garbage address, the HTTP client will fail immediately when it tries to connect.

func TestGetFilteredCountries_NetworkHttpError(t *testing.T) {

	_ = web.AppConfig.Set("restcountriesurl", "http://invalid-url-string-structure")

	// Grabbing a fresh instance of the service we want to test
	svc := &CountryService{}

	//running the function. The code inside will try to make a network request, hit our broken URL, and go right into the 'if err != nil' error
	results, err := svc.GetFilteredCountries("", "")

	//  expecting 'err' to be filled with the network error (not nil).
	assert.Error(t, err)

	// making sure 'results' is totally blank, because a broken connection shouldn't return corrupt data.
	assert.Nil(t, results)
}

// if external API server breaks and sends us complete garbage instead of proper country data.
func TestGetFilteredCountries_JsonParseError(t *testing.T) {

	// a temporary "fake" local server right inside our computer's memory.
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fake server responding with a successful 200 OK status...
		w.WriteHeader(http.StatusOK)

		// instead of returning a clean list of countries in JSON format,  it writes back a completely unparseable, scrambled text sentence.
		_, _ = w.Write([]byte("scrambled json payload string data bytes"))
	}))

	// Temporarily point the application's configuration to this fake local server's URL.
	_ = web.AppConfig.Set("restcountriesurl", server.URL)

	// Making sure this fake server shuts down cleanly when this specific test finishes.
	defer server.Close()

	svc := &CountryService{}

	// when  the code tries to parse that scrambled text sentence into Go data structures, it will fail.
	results, err := svc.GetFilteredCountries("", "")

	//  checking id  data-parsing safeguards worked perfectly.
	assert.Error(t, err)

	// checkinh if the results list empty instead of stuffing it with broken data
	assert.Nil(t, results)
}
