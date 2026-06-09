package services

import (
	"testing"
	"travelSphere/models"

	"github.com/stretchr/testify/assert"
)

// resetUserRegistry completely flushes our in-memory data store before each test runs which ensures old test data doesn't leak into new tests
func resetUserRegistry() {
	userRegistryLock.Lock()
	defer userRegistryLock.Unlock()
	userRegistry = make(map[string]*models.User)
	nextUserID = 1
}

// testing for AuthenticateUser with edge cases

func TestAuthenticateUser_CreateNewUser(t *testing.T) {

	// clearing out the registry and get a clean slate.
	resetUserRegistry()
	svc := &AuthService{}
	username := "sifa_khan"

	// trying  to authenticate someone who isn't registered yet.
	user, err := svc.AuthenticateUser(username)

	// the system should automatically reate a new profile, set their ID to "usr_1", and log a timestamp.
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "usr_1", user.ID)
	assert.Equal(t, username, user.Username)
	assert.NotZero(t, user.CreatedAt)

	// making sure the profile was safely saved and didn't just vanish in memory.
	userRegistryLock.RLock()
	savedUser, exists := userRegistry["usr_1"]
	userRegistryLock.RUnlock()

	assert.True(t, exists)
	assert.Equal(t, username, savedUser.Username)
}

func TestAuthenticateUser_ReturnExistingUser(t *testing.T) {
	// wipping  the data fresh.
	resetUserRegistry()
	svc := &AuthService{}
	username := "test_developer"

	// // logging the user in for the very first time whn they get created.
	firstUser, err1 := svc.AuthenticateUser(username)
	assert.NoError(t, err1)
	assert.NotNil(t, firstUser)
	assert.Equal(t, "usr_1", firstUser.ID)

	// trying logging in the exact same username again immediatly
	secondUser, err2 := svc.AuthenticateUser(username)

	// system , skip making a duplicate account, and return back their existing profile.
	assert.NoError(t, err2)
	assert.NotNil(t, secondUser)
	assert.Equal(t, firstUser.ID, secondUser.ID, "Should return the existing user instance matching the username")
	assert.Equal(t, firstUser.Username, secondUser.Username)
	assert.Equal(t, firstUser.CreatedAt, secondUser.CreatedAt)

	// ensuring auto-increment ID counter paused at 2 means it didn't generate a fake new ID.
	assert.Equal(t, 2, nextUserID)

}
