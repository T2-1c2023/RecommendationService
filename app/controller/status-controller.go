package controller

import (
	"net/http"

	dateGenerator "github.com/T2-1c2023/RecommendationService/app/utilities"
	"github.com/gin-gonic/gin"
)

var creationDate string = dateGenerator.GetCurrentDate()

type StatusController struct {
	CreationDate string
	Blocked      bool
}

type HealthResponse struct {
	CreationDate string `json:"creation_date"`
	LastResponse string `json:"last_response"`
	Description  string `json:"description"`
}

type ChangeStatusInput struct {
	Blocked bool `json:"blocked"`
}

func NewStatusController() StatusController {
	service := StatusController{
		CreationDate: dateGenerator.GetCurrentDate(),
		Blocked:      false,
	}
	return service
}

// GetStatus     godoc
// @Summary      Check the service's status.
// @Description  Returns a 200 code and JSON with a message.
// @Success      200
// @Router       / [get]
func (service *StatusController) GetStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Notifications microservice running correctly",
	})
}

// GetHealth     godoc
// @Summary      Get the service's health status.
// @Description  Returns a 200 code and JSON with the date the service started running and a description.
// @Success      200 {object} HealthResponse
// @Router       /health [get]
func (service *StatusController) GetHealth(c *gin.Context) {
	response := HealthResponse{
		CreationDate: creationDate,
		LastResponse: dateGenerator.GetCurrentDate(),
		Description:  "Recommendations microservice for FiuFit",
	}
	c.JSON(200, response)
}

// ChangeServiceStatus     	godoc
// @Summary      						Change the service's blocked status.
// @Description  						Changes the status of the service.
// @Param										user_info header string true "Decoded payload of the admin token"
// @Accept									json
// @Produce									json
// @Param										rule body ChangeStatusInput true "Blocked status of the service"
// @Success      						200 {object} ChangeStatusInput
// @Router       						/status [post]
func (service *StatusController) ChangeServiceStatus(c *gin.Context) {
	var input ChangeStatusInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	service.Blocked = input.Blocked
	c.JSON(http.StatusOK, input)
}

// GetServiceStatus		     	godoc
// @Summary      						Get the service's blocked status.
// @Description  						Get the status of the service.
// @Param										user_info header string true "Decoded payload of the admin token"
// @Produce									json
// @Success      						200 {object} ChangeStatusInput
// @Router       						/status [get]
func (service *StatusController) GetServiceStatus(c *gin.Context) {
	response := ChangeStatusInput{
		Blocked: service.Blocked,
	}
	c.JSON(http.StatusOK, response)
}

func (service *StatusController) ValidateBlockedStatus(c *gin.Context) {
	if service.Blocked {
		c.AbortWithStatusJSON(http.StatusLocked, gin.H{"message": "Service blocked"})
		return
	}
	c.Next()
}
