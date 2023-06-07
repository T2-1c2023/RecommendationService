package __test__

import (
	"encoding/json"

	mock "github.com/T2-1c2023/RecommendationService/__mock__"
	"github.com/T2-1c2023/RecommendationService/app/controller"
	"github.com/T2-1c2023/RecommendationService/app/model"
	"github.com/T2-1c2023/RecommendationService/app/routes"
	"github.com/gin-gonic/gin"
)

func setUpRouter(interestsRuleEnabled bool,
	proximityRuleEnabled bool) (*gin.Engine, []model.Training) {
	rulesRepositoryMock := mock.NewRulesRepositoryMock(
		interestsRuleEnabled,
		proximityRuleEnabled,
	)
	trainingServiceMock := mock.NewTrainingServiceMock()
	userServiceMock := mock.NewUserServiceMock()
	recommendationController := controller.RecommendationController{
		Repo:            &rulesRepositoryMock,
		UserService:     &userServiceMock,
		TrainingService: &trainingServiceMock,
	}
	proximityRuleController := controller.ProximityRuleController{
		Repo: &rulesRepositoryMock,
	}
	interestsRuleController := controller.InterestsRuleController{
		Repo: &rulesRepositoryMock,
	}
	router := routes.SetupRouter(&proximityRuleController,
		&interestsRuleController, &recommendationController)

	return router, trainingServiceMock.Trainings
}

func setUpErrorRouter() (*gin.Engine, []model.Training) {
	rulesRepositoryMock := mock.NewErrorRulesRepositoryMock()
	trainingServiceMock := mock.NewTrainingServiceMock()
	userServiceMock := mock.NewUserServiceMock()
	recommendationController := controller.RecommendationController{
		Repo:            &rulesRepositoryMock,
		UserService:     &userServiceMock,
		TrainingService: &trainingServiceMock,
	}
	proximityRuleController := controller.ProximityRuleController{
		Repo: &rulesRepositoryMock,
	}
	interestsRuleController := controller.InterestsRuleController{
		Repo: &rulesRepositoryMock,
	}
	router := routes.SetupRouter(&proximityRuleController,
		&interestsRuleController, &recommendationController)

	return router, trainingServiceMock.Trainings
}

func getUserInfo(isAdmin bool) string {
	userInfo, _ := json.Marshal(
		struct {
			Id      int  `json:"id"`
			IsAdmin bool `json:"is_admin"`
		}{
			Id:      1,
			IsAdmin: isAdmin,
		},
	)
	return string(userInfo)
}