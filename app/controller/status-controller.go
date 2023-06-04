package controller

import (
	dateGenerator "github.com/T2-1c2023/RecommendationService/app/utilities"
	"github.com/gin-gonic/gin"
)

var creationDate string = dateGenerator.GetCurrentDate()

// GetStatus     godoc
// @Summary      Check the service's status.
// @Description  Returns a 200 code and JSON with a message.
// @Success      200
// @Router       / [get]
func GetStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Notifications microservice running correctly",
	})
}

type HealthResponse struct {
	CreationDate string `json:"creation_date"`
	LastResponse string `json:"last_response"`
	Description  string `json:"description"`
}

// GetHealth     godoc
// @Summary      Get the service's health status.
// @Description  Returns a 200 code and JSON with the date the service started running and a description.
// @Success      200 {object} HealthResponse
// @Router       /health [get]
func GetHealth(c *gin.Context) {
	response := HealthResponse{
		CreationDate: creationDate,
		LastResponse: dateGenerator.GetCurrentDate(),
		Description:  "Recommendations microservice for FiuFit",
	}
	c.JSON(200, response)
}
