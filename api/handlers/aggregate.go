package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AggregateREsponse(c *gin.Context) {
	response1 := getService1Response()
	response2 := getService2Response()

	aggregateResponse := map[string]interface{}{
		"service1": response1,
		"service2": response2,
	}

	c.JSON(http.StatusOK, aggregateResponse)
}

func getService1Response() interface{} {
	return map[string]string{"message": "Service 1 response"}
}

func getService2Response() interface{} {
	return map[string]string{"message": "Service 2 response"}
}
