package controller

import (
	"fmt"
	"net/http"

	"github.com/T2-1c2023/RecommendationService/app/utilities"
	"github.com/gin-gonic/gin"
)

var creationDate string = utilities.GetCurrentDate()

type StatusController struct {
	CreationDate string
	Blocked      bool
	Logger       utilities.ILogger
}

type HealthResponse struct {
	CreationDate string `json:"creation_date"`
	LastResponse string `json:"last_response"`
	Description  string `json:"description"`
}

type ChangeStatusInput struct {
	Blocked bool `json:"blocked"`
}

// GetStatus     godoc
// @Summary      Check the service's status.
// @Description  Returns a 200 code and JSON with a message.
// @Success      200
// @Router       / [get]
func (controller *StatusController) GetStatus(c *gin.Context) {
	controller.Logger.LogInfo("Returning /")
	c.JSON(200, gin.H{
		"message": "Notifications microservice running correctly",
	})
}

// GetHealth     godoc
// @Summary      Get the service's health status.
// @Description  Returns a 200 code and JSON with the date the service started running and a description.
// @Success      200 {object} HealthResponse
// @Router       /health [get]
func (controller *StatusController) GetHealth(c *gin.Context) {
	response := HealthResponse{
		CreationDate: creationDate,
		LastResponse: utilities.GetCurrentDate(),
		Description:  "Recommendations microservice for FiuFit",
	}
	controller.Logger.LogInfo("Returning health response")
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
func (controller *StatusController) ChangeServiceStatus(c *gin.Context) {
	var input ChangeStatusInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		controller.Logger.LogError(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}
	controller.Blocked = input.Blocked
	controller.Logger.LogDebug(fmt.Sprintf("Blocked status: %t\n", controller.Blocked))
	c.JSON(http.StatusOK, input)
}

// GetServiceStatus		     	godoc
// @Summary      						Get the service's blocked status.
// @Description  						Get the status of the service.
// @Param										user_info header string true "Decoded payload of the admin token"
// @Produce									json
// @Success      						200 {object} ChangeStatusInput
// @Router       						/status [get]
func (controller *StatusController) GetServiceStatus(c *gin.Context) {
	response := ChangeStatusInput{
		Blocked: controller.Blocked,
	}
	controller.Logger.LogDebug(fmt.Sprintf("Blocked status: %t\n", controller.Blocked))
	c.JSON(http.StatusOK, response)
}

func (controller *StatusController) ValidateBlockedStatus(c *gin.Context) {
	if controller.Blocked {
		controller.Logger.LogDebug("Blocked incoming request")
		c.AbortWithStatusJSON(http.StatusLocked, gin.H{"message": "Service blocked"})
		return
	}
	c.Next()
}
