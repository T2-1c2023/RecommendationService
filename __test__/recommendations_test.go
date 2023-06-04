package __test__

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	mock "github.com/T2-1c2023/RecommendationService/__mock__"
	"github.com/T2-1c2023/RecommendationService/app/controller"
	"github.com/T2-1c2023/RecommendationService/app/model"
	routes "github.com/T2-1c2023/RecommendationService/app/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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

func sendRequest(router *gin.Engine) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodGet, "/recommended", nil)
	userInfo, _ := json.Marshal(
		struct {
			Id      int  `json:"id"`
			IsAdmin bool `json:"is_admin"`
		}{
			Id:      1,
			IsAdmin: true,
		})
	req.Header.Set("user_info", string(userInfo))

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	return recorder
}

func TestGetRecommendationsReturnsEmptyListWithRulesDisabled(t *testing.T) {
	router, _ := setUpRouter(false, false)
	recorder := sendRequest(router)

	assert.Equal(t, 200, recorder.Code)
}

func TestGetRecommendationsReturnsTrainingsWithProximityRuleEnabled(t *testing.T) {
	router, allTrainings := setUpRouter(false, true)
	recorder := sendRequest(router)

	var trainings []model.Training
	json.NewDecoder(recorder.Body).Decode(&trainings)

	assert.Equal(t, 200, recorder.Code)
	assert.ElementsMatch(t, allTrainings[:2], trainings)
}

func TestGetRecommendationsReturnsSomeTrainingsWithInterestsRuleEnabled(t *testing.T) {
	router, allTrainings := setUpRouter(true, false)
	recorder := sendRequest(router)

	var trainings []model.Training
	json.NewDecoder(recorder.Body).Decode(&trainings)

	assert.Equal(t, 200, recorder.Code)
	finalTrainings := []model.Training{allTrainings[0], allTrainings[2]}
	assert.ElementsMatch(t, finalTrainings, trainings)
}

func TestGetRecommendationsReturnsSomeTrainingsWithBothRulesEnabled(t *testing.T) {
	router, allTrainings := setUpRouter(true, true)
	recorder := sendRequest(router)

	var trainings []model.Training
	json.NewDecoder(recorder.Body).Decode(&trainings)

	assert.Equal(t, 200, recorder.Code)
	finalTrainings := []model.Training{allTrainings[0]}
	assert.ElementsMatch(t, finalTrainings, trainings)
}
