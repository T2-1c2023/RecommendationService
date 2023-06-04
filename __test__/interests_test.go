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

func TestUpdateInterestsRule(t *testing.T) {
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

	interestsRule := model.InterestsRule{
		Enabled: true,
	}
	payload, err := json.Marshal(interestsRule)
	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest(http.MethodPatch, "/rules/interests", bytes.NewBuffer(payload))

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

func TestGetInterestsRule(t *testing.T) {
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

	// Create a test HTTP request to the / endpoint
	req, _ := http.NewRequest(http.MethodGet, "/rules/interests", nil)

	userInfo, _ := json.Marshal(
		struct {
			Id      int  `json:"id"`
			IsAdmin bool `json:"is_admin"`
		}{
			Id:      1,
			IsAdmin: true,
		})
	req.Header.Set("user_info", string(userInfo))

	// Create a test HTTP recorder
	recorder := httptest.NewRecorder()

	// Serve the request and record the response
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}

func TestUpdateInterestsRuleError(t *testing.T) {
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

	interestsRule := model.InterestsRule{
		Enabled: true,
	}
	payload, err := json.Marshal(interestsRule)
	if err != nil {
		log.Fatal(err)
	}
	// Create a test HTTP request to the / endpoint
	req, _ := http.NewRequest(http.MethodPatch, "/rules/interests", bytes.NewBuffer(payload))

	userInfo, _ := json.Marshal(
		struct {
			Id      int  `json:"id"`
			IsAdmin bool `json:"is_admin"`
		}{
			Id:      1,
			IsAdmin: true,
		})
	req.Header.Set("user_info", string(userInfo))

	// Create a test HTTP recorder
	recorder := httptest.NewRecorder()

	// Serve the request and record the response
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 500, recorder.Code)
}

func TestGetInterestsRuleError(t *testing.T) {
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

	// Create a test HTTP request to the / endpoint
	req, _ := http.NewRequest(http.MethodGet, "/rules/interests", nil)

	userInfo, _ := json.Marshal(
		struct {
			Id      int  `json:"id"`
			IsAdmin bool `json:"is_admin"`
		}{
			Id:      1,
			IsAdmin: true,
		})
	req.Header.Set("user_info", string(userInfo))

	// Create a test HTTP recorder
	recorder := httptest.NewRecorder()

	// Serve the request and record the response
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 500, recorder.Code)
}
