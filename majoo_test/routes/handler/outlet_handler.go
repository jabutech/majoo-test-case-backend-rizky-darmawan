package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/majoo-test/models/web"
	"github.com/majoo-test/service"
)

type outletHandler struct {
	outletService   service.OutletService
	merchantService service.MerchantService
}

func NewOutletHandler(outletService service.OutletService, merchantService service.MerchantService) *outletHandler {
	return &outletHandler{outletService, merchantService}
}

func (h *outletHandler) GetListOutletTransactions(c *gin.Context) {
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

	// Get outlet id from uri
	var outletID web.OutletID
	err := c.ShouldBindUri(&outletID)
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

	// Find merchant by user id is current login
	merchant, _ := h.merchantService.FindByUserID(currentUserID)

	// Find outlet by id
	outlet, _ := h.outletService.FindByID(outletID)

	if merchant.ID != outlet.ID {
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

	// Find outlet transactions
	outletTransactions, _ := h.outletService.GetListOutletTransactions(outletID, &pagination, dateFrom, dateTo)
	outletFormatter := web.FormatOutletTransactions(outletTransactions)

	meta := web.Meta{
		PerPage:     perPage,
		CurrentPage: currentPage,
	}

	// Create format response
	response := web.ApiResponseWithDataPagination(
		http.StatusOK,
		"success",
		"List of outlet",
		outletFormatter,
		meta,
	)
	c.JSON(http.StatusOK, response)
}
