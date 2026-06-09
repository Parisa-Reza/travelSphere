package services

import (
	"fmt"
	"strings"
	"sync"
	"time"
	"travelSphere/models"
)

type WishlistService struct{}

var (
	wishlistStore = make(map[string]*models.WishlistEntry)
	wishlistLock  sync.RWMutex
	nextEntryID   = 1
)

func (s *WishlistService) GetUserWishlist(userID string) []models.WishlistEntry {
	wishlistLock.RLock()
	defer wishlistLock.RUnlock()

	var items []models.WishlistEntry
	for _, entry := range wishlistStore {
		if entry.UserID == userID {
			items = append(items, *entry)
		}
	}
	return items
}

func (s *WishlistService) AddToWishlist(userID, countryName string) (*models.WishlistEntry, error) {
	wishlistLock.Lock()
	defer wishlistLock.Unlock()

	slug := strings.ToLower(strings.ReplaceAll(countryName, " ", "-"))

	for _, entry := range wishlistStore {
		if entry.UserID == userID && entry.Slug == slug {
			return nil, fmt.Errorf("this country is already in your wishlist")
		}
	}

	newID := fmt.Sprintf("wl_%d", nextEntryID)
	nextEntryID++

	newEntry := &models.WishlistEntry{
		ID:          newID,
		UserID:      userID,
		CountryName: countryName,
		Slug:        slug,
		Note:        "",
		Status:      "Planned",
		CreatedAt:   time.Now(),
	}

	wishlistStore[newID] = newEntry
	return newEntry, nil
}

func (s *WishlistService) UpdateWishlist(id, userID, note, status string) (*models.WishlistEntry, error) {
	wishlistLock.Lock()
	defer wishlistLock.Unlock()

	entry, exists := wishlistStore[id]
	if !exists || entry.UserID != userID {
		return nil, fmt.Errorf("entry not found or unauthorized")
	}

	if status != "Planned" && status != "Visited" {
		return nil, fmt.Errorf("invalid status value")
	}

	entry.Note = note
	entry.Status = status
	return entry, nil
}

func (s *WishlistService) DeleteWishlist(id, userID string) error {
	wishlistLock.Lock()
	defer wishlistLock.Unlock()

	entry, exists := wishlistStore[id]
	if !exists || entry.UserID != userID {
		return fmt.Errorf("entry not found or unauthorized")
	}

	delete(wishlistStore, id)
	return nil
}
