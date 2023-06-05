package __test__

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	mock "github.com/T2-1c2023/RecommendationService/__mock__"
	"github.com/T2-1c2023/RecommendationService/app/controller"
	"github.com/T2-1c2023/RecommendationService/app/model"
	routes "github.com/T2-1c2023/RecommendationService/app/routes"
	"github.com/stretchr/testify/assert"
)

func TestUpdateProximityRule(t *testing.T) {
	rulesRepositoryMock := mock.NewRulesRepositoryMock(false, false)
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

	proximityRule := model.ProximityRule{
		Radius:  3,
		Enabled: true,
	}
	payload, err := json.Marshal(proximityRule)
	if err != nil {
		log.Fatal(err)
	}
	req, _ := http.NewRequest(http.MethodPatch, "/recommended/rules/proximity", bytes.NewBuffer(payload))

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

	assert.Equal(t, 201, recorder.Code)
}

func TestGetProximityRule(t *testing.T) {
	rulesRepositoryMock := mock.NewRulesRepositoryMock(false, false)
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

	req, _ := http.NewRequest(http.MethodGet, "/recommended/rules/proximity", nil)

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

	assert.Equal(t, 200, recorder.Code)
}

func TestUpdateProximityRuleError(t *testing.T) {
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

	proximityRule := model.ProximityRule{
		Radius:  3,
		Enabled: true,
	}
	payload, err := json.Marshal(proximityRule)
	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest(http.MethodPatch, "/recommended/rules/proximity", bytes.NewBuffer(payload))

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

	assert.Equal(t, 500, recorder.Code)
}

func TestGetProximityRuleError(t *testing.T) {
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

	req, _ := http.NewRequest(http.MethodGet, "/recommended/rules/proximity", nil)

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

	assert.Equal(t, 500, recorder.Code)
}
