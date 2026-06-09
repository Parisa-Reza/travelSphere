package services

import (
	"travelSphere/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
	"github.com/beego/beego/v2/server/web"
)
	


type AttractionService struct{}

func (s *AttractionService) GetPopularAttractions() ([]models.Attraction, error) {
	

	// Extract the secure string variable from  conf/app.conf file
	apiKey, err := web.AppConfig.String("opentripmapkey")
	if err != nil || apiKey == "" {
		return nil, fmt.Errorf("failed to retrieve opentripmapkey from app.conf configuration")
	}

     baseSvcUrl, err := web.AppConfig.String("opentripmapurl")


	//  Fetch the secure API key string from your app.conf file
	apiKey, error := web.AppConfig.String("opentripmapkey") 
	if error != nil || apiKey == "" {
		return nil, fmt.Errorf("failed to retrieve opentripmapkey from app.conf")
	}

	//  Format the endpoint dynamically using two %s 
	apiUrl := fmt.Sprintf("%s/radius?radius=500&lat=48.8584&lon=2.2945&kinds=interesting_places&rate=3&limit=5&format=json&apikey=%s", baseSvcUrl, apiKey)



	client := &http.Client{Timeout: 6 * time.Second}
	resp, err := client.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("attractions API returned bad status: %d", resp.StatusCode)
	}

	var rawList []models.Attraction
	if err := json.NewDecoder(resp.Body).Decode(&rawList); err != nil {
		return nil, err
	}

	var finalizedList []models.Attraction
	seenNames := make(map[string]bool)

	for _, item := range rawList {
		trimmedName := strings.TrimSpace(item.Name)
		
		
		if trimmedName == "" || seenNames[strings.ToLower(trimmedName)] {
			continue
		}

		
		cleanKinds := strings.ReplaceAll(item.Kinds, "_", " ")
		kindsArray := strings.Split(cleanKinds, ",")
		var shortKinds []string
		
		// Grab up to the first two category descriptors to keep the UI clean
		for i, k := range kindsArray {
			if i >= 2 {
				break
			}
			shortKinds = append(shortKinds, strings.TrimSpace(k))
		}
		item.DisplayKinds = strings.Join(shortKinds, ", ")

		seenNames[strings.ToLower(trimmedName)] = true
		finalizedList = append(finalizedList, item)
	}

	return finalizedList, nil
}