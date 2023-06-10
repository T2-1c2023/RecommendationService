package controller

import (
	"net/http"

	"github.com/T2-1c2023/RecommendationService/app/model"
	"github.com/T2-1c2023/RecommendationService/app/persistence"
	"github.com/T2-1c2023/RecommendationService/app/utilities"
	"github.com/gin-gonic/gin"
)

type ProximityRuleController struct {
	Repo   persistence.IRulesRepository
	Logger utilities.ILogger
}

// ModifyProximityRule     	godoc
// @Summary      						Update the maximum radius for recommended trainings.
// @Description  						Update the maximum radius for recommended trainings.
// @Tags										Proximity Rule
// @Param										user_info header string true "Proximity rule changes"
// @Accept									json
// @Produce									json
// @Param										rule body model.ProximityRule true "Proximity rule changes"
// @Success      						201
// @Failure									400
// @Failure									423
// @Failure									500
// @Router       						/recommended/rules/proximity [patch]
func (controller *ProximityRuleController) ModifyProximityRule(c *gin.Context) {
	var input model.ProximityRule
	err := c.ShouldBindJSON(&input)
	if err != nil {
		controller.Logger.LogInfo("Bad request, returning 400")
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = controller.Repo.UpdateProximityRule(&input)
	if err != nil {
		controller.Logger.LogError(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	controller.Logger.LogInfo("Rule updated")
	c.JSON(http.StatusCreated, gin.H{"message": "Rule updated"})
}

// GetProximityRule     		godoc
// @Summary      						Get the current settings of the proximity rule.
// @Description  						Get the current settings of the proximity rule.
// @Tags										Proximity Rule
// @Param										user_info header string true "Proximity rule changes"
// @Produce									json
// @Success      						200 {object} model.ProximityRule
// @Failure									500
// @Router       						/recommended/rules/proximity [get]
func (controller *ProximityRuleController) GetProximityRule(c *gin.Context) {
	rule, err := controller.Repo.GetProximityRule()
	if err != nil {
		controller.Logger.LogError(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	controller.Logger.LogInfo("Returning proximity rule")
	c.JSON(http.StatusOK, rule)
}
