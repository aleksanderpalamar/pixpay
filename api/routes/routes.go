package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/aleksanderpalamar/pixpay/api/handlers"
	"github.com/aleksanderpalamar/pixpay/api/middleware"
)

func SetupRouter(r *gin.Engine) {
	r.Use(middleware.RateLimitMiddleware())
	r.Use(middleware.CacheMiddleware())
	r.Use(middleware.LoggerMiddleware())

	r.GET("/payments", handlers.GetPayments)
	r.POST("/payments", handlers.CreatePayment)
	r.GET("/payments/:id", handlers.GetPaymentByID)
	r.PUT("/payments/:id", handlers.UpdatePayment)
	r.DELETE("/payments/:id", handlers.DeletePayment)
}
