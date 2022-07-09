package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/majoo-test/middleware"
	"github.com/majoo-test/repository"
	"github.com/majoo-test/routes/handler"
	"github.com/majoo-test/service"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	// Repository
	userRepository := repository.NewRepositoryUser(db)
	merchantRepository := repository.NewRepositoryMerchant(db)
	outletRepository := repository.NewRepositoryOutlet(db)

	// Service
	userService := service.NewServiceUser(userRepository)
	merchantService := service.NewMerchantService(merchantRepository)
	outletService := service.NewOutletService(outletRepository)

	// Handler
	authHandler := handler.NewAuthHandler(userService)
	merchantHandler := handler.NewMerchantHandler(merchantService)
	outletHandler := handler.NewOutletHandler(outletService, merchantService)

	router := gin.Default()
	router.Use(cors.Default())

	api := router.Group("/api/v1")

	// Endpoint auth
	auth := api.Group("/auth")
	auth.POST("/login", authHandler.Login)

	// Endpoint merchants
	api.GET("/merchants/:merchant_id/reports", middleware.AuthMiddleware(), merchantHandler.GetListMerchantTransactions)

	// Endpoint outlet
	api.GET("/outlets/:outlet_id/reports", middleware.AuthMiddleware(), outletHandler.GetListOutletTransactions)
	return router
}
