package controller

import (
	"net/http"

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
	userInfo string) ([]model.Training, error) {
	var trainings []model.Training
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
				)
			if err != nil {
				return trainings, err
			}
			trainings = append(trainings, someTrainings...)
		}
	} else {
		trainings, err =
			controller.TrainingService.GetAllTrainings(userInfo)
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

// GetRecommendations     	godoc
// @Summary      						Get recommended trainings according to current ruleset.
// @Description  						Get recommended trainings according to current ruleset.
// @Tags										Recommendations
// @Param										user_info header string true "Proximity rule changes"
// @Produce									json
// @Success      						200 {array} model.Training
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

	trainings, err := controller.getTrainings(proximityRule, c.GetHeader("user_info"))
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusServiceUnavailable,
			gin.H{"message": err.Error()},
		)
		return
	}

	if !interestsRule.Enabled {
		c.JSON(http.StatusOK, trainings)
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

	c.JSON(http.StatusOK, filteredTrainings)
}
