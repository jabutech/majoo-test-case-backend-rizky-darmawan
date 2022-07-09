package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/majoo-test/models/web"
	"github.com/majoo-test/service"
)

type merchantHandler struct {
	merchantService service.MerchantService
}

func NewMerchantHandler(merchantService service.MerchantService) *merchantHandler {
	return &merchantHandler{merchantService}
}

func (h *merchantHandler) GetListMerchantTransactions(c *gin.Context) {
	perPage := 5                   // Default display data per page
	currentPage := 1               // Default current page displaying
	query := c.Request.URL.Query() // Request query

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "per_page":
			perPage, _ = strconv.Atoi(queryValue)
			break
		case "current_page":
			currentPage, _ = strconv.Atoi(queryValue)
			break
		}
	}

	pagination := web.Pagination{
		Limit: perPage,
		Page:  currentPage,
	}

	// Get merchant id from uri
	var merchantID web.MerchantID
	err := c.ShouldBindUri(&merchantID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := web.ApiResponseWithData(
			http.StatusBadRequest,
			"error",
			"Internal Server error",
			errorMessage,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Get data current user is logged in
	currentUser := c.MustGet("currentUser").(string)
	currentUserID, _ := strconv.Atoi(currentUser)

	// Find merchant by id
	merchant, _ := h.merchantService.FindByID(merchantID)
	if int64(currentUserID) != int64(merchant.UserID) {
		response := web.ApiResponseWithoutData(
			http.StatusForbidden,
			"error",
			"User not allowed",
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Query date from and date to for filter transaction
	dateFrom := c.Query("date_from")
	dateTo := c.Query("date_to")

	// Find merchants
	merchants, _ := h.merchantService.GetListMerchantTransactions(merchantID, &pagination, dateFrom, dateTo)
	merchantFormatter := web.FormatMerchantTransactions(merchants)

	meta := web.Meta{
		PerPage:     perPage,
		CurrentPage: currentPage,
	}

	// Create format response
	response := web.ApiResponseWithDataPagination(
		http.StatusOK,
		"success",
		"List of merchants",
		merchantFormatter,
		meta,
	)
	c.JSON(http.StatusOK, response)
}
