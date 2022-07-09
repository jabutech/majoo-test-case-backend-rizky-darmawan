package test

import (
	"log"
	"testing"

	"github.com/majoo-test/models/web"
	"github.com/majoo-test/repository"
	"github.com/majoo-test/service"
	"github.com/majoo-test/util"
	"github.com/stretchr/testify/assert"
)

// Test login success
func TestLoginSuccess(t *testing.T) {
	// Open Connection
	db := util.SetupDB()
	// Repository
	userRepository := repository.NewRepositoryUser(db)
	// Service
	userService := service.NewServiceUser(userRepository)

	// Payload
	payload := web.LoginRequest{
		UserName: "admin1",
		Password: "admin1",
	}

	// Login
	userLogin, err := userService.Login(payload)
	if err != nil {
		log.Panic(err)
	}

	assert.Equal(t, int64(1), userLogin.ID)
	assert.Equal(t, "Admin 1", userLogin.Name)
	assert.Equal(t, "admin1", userLogin.UserName)
	assert.Equal(t, int64(1), userLogin.CreatedBy)
	assert.Equal(t, int64(1), userLogin.UpdatedBy)

	assert.NotEmpty(t, userLogin.Password)
	assert.NotEmpty(t, userLogin.CreatedAt)
	assert.NotEmpty(t, userLogin.UpdatedAt)
}

// Test login failed
func TestLoginFailed(t *testing.T) {

	db := util.SetupDB()
	// Repository
	userRepository := repository.NewRepositoryUser(db)
	// Service
	userService := service.NewServiceUser(userRepository)
	t.Run("password wrong", func(t *testing.T) {
		// Payload
		payload := web.LoginRequest{
			UserName: "admin1",
			Password: "wrong",
		}

		// Login
		_, err := userService.Login(payload)
		assert.NotNil(t, err)
		assert.Equal(t, "email or password incorrect", err.Error())
	})

	t.Run("username wrong", func(t *testing.T) {
		// Payload
		payload := web.LoginRequest{
			UserName: "wrong",
			Password: "admin1",
		}

		// Login
		_, err := userService.Login(payload)
		assert.NotNil(t, err)
		assert.Equal(t, "email or password incorrect", err.Error())
	})

}

// Test Generate token success
func TestGenerateToken(t *testing.T) {
	db := util.SetupDB()
	// Repository
	userRepository := repository.NewRepositoryUser(db)
	// find by username
	user, err := userRepository.FindByUserName("admin1")
	if err != nil {
		log.Panic(err)
	}

	// Service
	userService := service.NewServiceUser(userRepository)

	// Login
	token, err := userService.GenerateToken(user)
	if err != nil {
		log.Panic(err)
	}
	assert.NotEmpty(t, token)
}
