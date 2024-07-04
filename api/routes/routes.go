package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/aleksanderpalamar/pixpay/api/handlers"
	"github.com/aleksanderpalamar/pixpay/api/middleware"
)

func SetuoRouter() *gin.Engine {
	router := gin.Default()

	// Middleware
	router.Use(middleware.AuthMiddleware())

	// Routes
	router.GET("/payments", handlers.GetPayments)
	router.POST("/payments", handlers.CreatePayment)
	router.GET("/payments/:id", handlers.GetPaymentByID)
	router.PUT("/payments/:id", handlers.UpdatePayment)
	router.DELETE("/payments/:id", handlers.DeletePayment)

	return router
}
