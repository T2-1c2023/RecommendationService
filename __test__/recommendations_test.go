package __test__

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/T2-1c2023/RecommendationService/app/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func sendRecommendedRequest(router *gin.Engine, isAdmin bool) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodGet, "/recommended", nil)
	userInfo := getUserInfo(isAdmin)
	req.Header.Set("user_info", userInfo)

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	return recorder
}

func TestGetRecommendationsReturnsEmptyListWithRulesDisabled(t *testing.T) {
	router, _ := setUpRouter(false, false)
	recorder := sendRecommendedRequest(router, false)

	assert.Equal(t, 200, recorder.Code)
}

func TestGetRecommendationsReturnsTrainingsWithProximityRuleEnabled(t *testing.T) {
	router, allTrainings := setUpRouter(false, true)
	recorder := sendRecommendedRequest(router, false)

	var trainings []model.Training
	json.NewDecoder(recorder.Body).Decode(&trainings)

	assert.Equal(t, 200, recorder.Code)
	assert.ElementsMatch(t, allTrainings[:2], trainings)
}

func TestGetRecommendationsReturnsSomeTrainingsWithInterestsRuleEnabled(t *testing.T) {
	router, allTrainings := setUpRouter(true, false)
	recorder := sendRecommendedRequest(router, false)

	var trainings []model.Training
	json.NewDecoder(recorder.Body).Decode(&trainings)

	assert.Equal(t, 200, recorder.Code)
	finalTrainings := []model.Training{allTrainings[0], allTrainings[2]}
	assert.ElementsMatch(t, finalTrainings, trainings)
}
