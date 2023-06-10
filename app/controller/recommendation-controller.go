package controller

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/T2-1c2023/RecommendationService/app/model"
	"github.com/T2-1c2023/RecommendationService/app/persistence"
	"github.com/T2-1c2023/RecommendationService/app/services"
	"github.com/T2-1c2023/RecommendationService/app/utilities"
	"github.com/gin-gonic/gin"
)

type RecommendationController struct {
	Repo            persistence.IRulesRepository
	UserService     services.IUserService
	TrainingService services.ITrainingService
	Logger          utilities.ILogger
}

func (controller *RecommendationController) getTrainings(proximityRule model.ProximityRule,
	userInfo string, queryParams map[string]string) ([]model.Training, error) {
	trainings := []model.Training{}
	var err error
	if proximityRule.Enabled {
		var closeTrainers []model.User
		closeTrainers, err =
			controller.UserService.GetAllTrainersInProximity(
				proximityRule.Radius,
				userInfo,
			)
		if err != nil {
			controller.Logger.LogError(err)
			return trainings, err
		}
		for _, trainer := range closeTrainers {
			someTrainings, err :=
				controller.TrainingService.GetTrainingsFromTrainerId(
					trainer.Id,
					userInfo,
					queryParams,
				)
			if err != nil {
				controller.Logger.LogError(err)
				return trainings, err
			}
			trainings = append(trainings, someTrainings...)
		}
	} else {
		trainings, err =
			controller.TrainingService.GetAllTrainings(userInfo, queryParams)
		if err != nil {
			controller.Logger.LogError(err)
			return trainings, err
		}
	}
	return trainings, nil
}

func (controller *RecommendationController) getInterests(userInfo string) ([]model.Interest, error) {
	var interests []model.Interest
	userInfoObject, err := model.NewUserFromUserInfo(userInfo)
	if err != nil {
		controller.Logger.LogError(err)
		return interests, err
	}
	user, err := controller.UserService.GetUserById(userInfoObject.Id, userInfo)
	if err != nil {
		controller.Logger.LogError(err)
		return interests, err
	}
	return user.Interests, nil
}

func (controller *RecommendationController) filterTrainingsByInterests(trainings []model.Training,
	interests []model.Interest) []model.Training {
	var filteredTrainings []model.Training
	for _, training := range trainings {
		for _, interest := range interests {
			if training.Type == interest.Description {
				filteredTrainings = append(filteredTrainings, training)
			}
		}
	}
	return filteredTrainings
}

func (controller *RecommendationController) getQueryParams(c *gin.Context) map[string]string {
	queryParams := make(map[string]string)
	paramsString := os.Getenv("TRAININGS_QUERY_PARAMS")
	params := strings.Split(paramsString, ",")
	for _, param := range params {
		value := c.Query(param)
		if value != "" {
			queryParams[param] = value
		}
	}
	return queryParams
}

// GetRecommendations     	godoc
// @Summary      						Get recommended trainings according to current ruleset.
// @Description  						Get recommended trainings according to current ruleset.
// @Tags										Recommendations
// @Param										user_info header string true "Proximity rule changes"
// @Param 									title query string false "Title of the training"
// @Param 									trainer_id query int false "ID of the training owner"
// @Param 									type_id query int false "ID of the training type"
// @Param 									severity query int false "Severity of the training"
// @Param 									blocked query bool false "Whether the training is blocked"
// @Produce									json
// @Success      						200 {array} model.Training
// @Failure									423
// @Failure									500
// @Failure									503
// @Router       						/recommended [get]
func (controller *RecommendationController) GetRecommendations(c *gin.Context) {
	controller.Logger.LogInfo("GET /recommended")
	proximityRule, err1 := controller.Repo.GetProximityRule()
	interestsRule, err2 := controller.Repo.GetInterestsRule()
	if err1 != nil || err2 != nil {
		controller.Logger.LogError(fmt.Errorf("database error"))
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": "Database error"})
	}
	controller.Logger.LogDebug(fmt.Sprintf("Proximity rule enabled: %t", proximityRule.Enabled))
	controller.Logger.LogDebug(fmt.Sprintf("Interests rule enabled: %t", interestsRule.Enabled))

	if !proximityRule.Enabled && !interestsRule.Enabled {
		controller.Logger.LogWarn("Neither rule is enabled, returning empty list")
		trainings := []model.Training{}
		c.JSON(http.StatusOK, trainings)
		return
	}

	queryParams := controller.getQueryParams(c)
	controller.Logger.LogDebug("User Info: " + c.GetHeader("user_info"))
	trainings, err := controller.getTrainings(
		proximityRule,
		c.GetHeader("user_info"),
		queryParams,
	)
	if err != nil {
		controller.Logger.LogError(err)
		c.AbortWithStatusJSON(
			http.StatusServiceUnavailable,
			gin.H{"message": err.Error()},
		)
		return
	}

	controller.Logger.LogDebug(fmt.Sprintf("Amount of trainings retrieved: %d", len(trainings)))

	if !interestsRule.Enabled {
		if len(trainings) == 0 {
			controller.Logger.LogInfo("Returning empty training list")
			c.JSON(http.StatusOK, make([]model.Training, 0))
		} else {
			controller.Logger.LogInfo("Returning training list")
			c.JSON(http.StatusOK, trainings)
		}
		return
	}

	interests, err := controller.getInterests(c.GetHeader("user_info"))
	if err != nil {
		controller.Logger.LogError(err)
		c.AbortWithStatusJSON(
			http.StatusServiceUnavailable,
			gin.H{"message": err.Error()},
		)
		return
	}

	filteredTrainings := controller.filterTrainingsByInterests(trainings, interests)
	controller.Logger.LogDebug(fmt.Sprintf("Amount of interests retrieved: %d", len(interests)))
	controller.Logger.LogDebug(fmt.Sprintf("Amount of filtered trainings: %d", len(filteredTrainings)))
	if len(filteredTrainings) == 0 {
		controller.Logger.LogInfo("Returning empty training list")
		c.JSON(http.StatusOK, make([]model.Training, 0))
	} else {
		controller.Logger.LogInfo("Returning training list")
		c.JSON(http.StatusOK, filteredTrainings)
	}
}
