package test

import (
	"log"
	"testing"

	"github.com/majoo-test/repository"
	"github.com/majoo-test/util"
	"github.com/stretchr/testify/assert"
)

func TestFindByUsername(t *testing.T) {
	db := util.SetupDB()
	userRepository := repository.NewRepositoryUser(db)
	t.Run("find admin 1", func(t *testing.T) {
		// Test find by username
		user, err := userRepository.FindByUserName("admin1")
		if err != nil {
			log.Panic(err)
		}

		assert.Equal(t, int64(1), user.ID)
		assert.Equal(t, "Admin 1", user.Name)
		assert.Equal(t, "admin1", user.UserName)
		assert.Equal(t, int64(1), user.CreatedBy)
		assert.Equal(t, int64(1), user.UpdatedBy)

		assert.NotEmpty(t, user.Password)
		assert.NotEmpty(t, user.CreatedAt)
		assert.NotEmpty(t, user.UpdatedAt)
	})
	t.Run("find admin 1", func(t *testing.T) {
		// Test find by username
		user, err := userRepository.FindByUserName("admin2")
		if err != nil {
			log.Panic(err)
		}

		assert.Equal(t, int64(2), user.ID)
		assert.Equal(t, "Admin 2", user.Name)
		assert.Equal(t, "admin2", user.UserName)
		assert.Equal(t, int64(2), user.CreatedBy)
		assert.Equal(t, int64(2), user.UpdatedBy)

		assert.NotEmpty(t, user.Password)
		assert.NotEmpty(t, user.CreatedAt)
		assert.NotEmpty(t, user.UpdatedAt)
	})

}
