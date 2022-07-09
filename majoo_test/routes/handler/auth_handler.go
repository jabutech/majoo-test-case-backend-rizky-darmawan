package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/majoo-test/models/web"
	"github.com/majoo-test/service"
)

type authHandler struct {
	userService service.UserService
}

func NewAuthHandler(userService service.UserService) *authHandler {
	return &authHandler{userService}
}

func (h *authHandler) Login(c *gin.Context) {
	var req web.LoginRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		errors := web.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := web.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Login failed",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Login
	userLogin, err := h.userService.Login(req)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := web.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Login failed",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Generate token
	token, err := h.userService.GenerateToken(userLogin)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := web.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Login failed",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	fieldToken := gin.H{"token": token}
	// Create format response
	response := web.ApiResponseWithData(
		http.StatusOK,
		"success",
		"You have successfully Login",
		fieldToken,
	)
	c.JSON(http.StatusOK, response)
}
