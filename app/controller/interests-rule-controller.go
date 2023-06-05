package controller

import (
	"net/http"

	"github.com/T2-1c2023/RecommendationService/app/model"
	"github.com/T2-1c2023/RecommendationService/app/persistence"
	"github.com/gin-gonic/gin"
)

type InterestsRuleController struct {
	Repo persistence.IRulesRepository
}

// ModifyInterestsRule     	godoc
// @Summary      						Update the interests rule' status for recommended trainings.
// @Description  						Update the interests rule' status for recommended trainings.
// @Tags										Interests Rule
// @Param										user_info header string true "Proximity rule changes"
// @Accept									json
// @Produce									json
// @Param										rule body model.InterestsRule true "Interests rule changes"
// @Success      						201
// @Failure									400
// @Failure									500
// @Router       						/recommended/rules/interests [patch]
func (controller *InterestsRuleController) ModifyInterestsRule(c *gin.Context) {
	var input model.InterestsRule
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = controller.Repo.UpdateInterestsRule(&input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Rule updated"})
}

// GetInterestsRule     		godoc
// @Summary      						Get the current settings of the interests rule.
// @Description  						Get the current settings of the interests rule.
// @Tags										Interests Rule
// @Param										user_info header string true "Proximity rule changes"
// @Produce									json
// @Success      						200 {object} model.InterestsRule
// @Failure									500
// @Router       						/recommended/rules/interests [get]
func (controller *InterestsRuleController) GetInterestsRule(c *gin.Context) {
	rule, err := controller.Repo.GetInterestsRule()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rule)
}
