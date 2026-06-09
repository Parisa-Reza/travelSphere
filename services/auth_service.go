package services

import (
	"fmt"
	"sync"
	"time"
	"travelSphere/models"
)

type AuthService struct{}

var (
	// Thread-safe map memory cache layer mapping unique string IDs to User entities
	userRegistry     = make(map[string]*models.User)
	userRegistryLock sync.RWMutex
	nextUserID       = 1
)

func (s *AuthService) AuthenticateUser(username string) (*models.User, error) {
	userRegistryLock.Lock()
	defer userRegistryLock.Unlock()

	// Check if this user name is already registered in memory
	for _, existingUser := range userRegistry {
		if existingUser.Username == username {
			return existingUser, nil
		}
	}

	// Multi-user system execution: Create a uniquely identified entity mapping
	newID := fmt.Sprintf("usr_%d", nextUserID)
	nextUserID++

	newUser := &models.User{
		ID:        newID,
		Username:  username,
		CreatedAt: time.Now(),
	}

	userRegistry[newID] = newUser
	return newUser, nil
}
