package controller

import (
	"net/http"
	"os"
	"strings"

	"github.com/T2-1c2023/RecommendationService/app/model"
	"github.com/T2-1c2023/RecommendationService/app/persistence"
	"github.com/T2-1c2023/RecommendationService/app/services"
	"github.com/gin-gonic/gin"
)

type RecommendationController struct {
	Repo            persistence.IRulesRepository
	UserService     services.IUserService
	TrainingService services.ITrainingService
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
				return trainings, err
			}
			trainings = append(trainings, someTrainings...)
		}
	} else {
		trainings, err =
			controller.TrainingService.GetAllTrainings(userInfo, queryParams)
		if err != nil {
			return trainings, err
		}
	}
	return trainings, nil
}

func (controller *RecommendationController) getInterests(userInfo string) ([]model.Interest, error) {
	var interests []model.Interest
	userInfoObject, err := model.NewUserFromUserInfo(userInfo)
	if err != nil {
		return interests, err
	}
	user, err := controller.UserService.GetUserById(userInfoObject.Id, userInfo)
	if err != nil {
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
	proximityRule, err1 := controller.Repo.GetProximityRule()
	interestsRule, err2 := controller.Repo.GetInterestsRule()
	if err1 != nil || err2 != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": "Database error"})
	}

	if !proximityRule.Enabled && !interestsRule.Enabled {
		trainings := []model.Training{}
		c.JSON(http.StatusOK, trainings)
		return
	}

	queryParams := controller.getQueryParams(c)
	trainings, err := controller.getTrainings(
		proximityRule,
		c.GetHeader("user_info"),
		queryParams,
	)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusServiceUnavailable,
			gin.H{"message": err.Error()},
		)
		return
	}

	if !interestsRule.Enabled {
		if len(trainings) == 0 {
			c.JSON(http.StatusOK, make([]model.Training, 0))
		} else {
			c.JSON(http.StatusOK, trainings)
		}
		return
	}

	interests, err := controller.getInterests(c.GetHeader("user_info"))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusServiceUnavailable,
			gin.H{"message": err.Error()},
		)
		return
	}

	filteredTrainings := controller.filterTrainingsByInterests(trainings, interests)
	if len(filteredTrainings) == 0 {
		c.JSON(http.StatusOK, make([]model.Training, 0))
	} else {
		c.JSON(http.StatusOK, filteredTrainings)
	}
}
