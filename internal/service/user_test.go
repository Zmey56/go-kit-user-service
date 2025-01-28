package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserService(t *testing.T) {
	// Initializing the service
	svc := NewUserService()

	// Testing CreateUser
	t.Run("CreateUser", func(t *testing.T) {
		user, err := svc.CreateUser("John Doe", "john.doe@example.com", 30)
		assert.NoError(t, err)
		assert.Equal(t, "John Doe", user.Name)
		assert.Equal(t, "john.doe@example.com", user.Email)
		assert.Equal(t, 30, user.Age)
	})

	// Testing GetUserByID
	t.Run("GetUserByID", func(t *testing.T) {
		user, _ := svc.CreateUser("Jane Doe", "jane.doe@example.com", 25)
		retrievedUser, err := svc.GetUserByID(user.ID)
		assert.NoError(t, err)
		assert.Equal(t, user, retrievedUser)
	})

	// Testing UpdateUser
	t.Run("UpdateUser", func(t *testing.T) {
		user, _ := svc.CreateUser("Old Name", "old@example.com", 40)
		updatedUser, err := svc.UpdateUser(user.ID, "New Name", "new@example.com", 41)
		assert.NoError(t, err)
		assert.Equal(t, "New Name", updatedUser.Name)
		assert.Equal(t, "new@example.com", updatedUser.Email)
		assert.Equal(t, 41, updatedUser.Age)
	})

	// Testing DeleteUser
	t.Run("DeleteUser", func(t *testing.T) {
		user, _ := svc.CreateUser("To Delete", "delete@example.com", 50)
		err := svc.DeleteUser(user.ID)
		assert.NoError(t, err)

		_, err = svc.GetUserByID(user.ID)
		assert.Error(t, err)
	})
}
