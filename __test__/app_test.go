// __test__/app_test.go
package __test__

import (
	"net/http"
	"net/http/httptest"
	"testing"

	mock "github.com/T2-1c2023/RecommendationService/__mock__"
	"github.com/T2-1c2023/RecommendationService/app/controller"
	routes "github.com/T2-1c2023/RecommendationService/app/routes"
	"github.com/stretchr/testify/assert"
)

func TestGetStatus(t *testing.T) {
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
	req, _ := http.NewRequest(http.MethodGet, "/", nil)

	// Create a test HTTP recorder
	recorder := httptest.NewRecorder()

	// Serve the request and record the response
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}
