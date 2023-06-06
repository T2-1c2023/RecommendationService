package __test__

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/T2-1c2023/RecommendationService/app/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func sendPatchProximityRequest(router *gin.Engine,
	proximityRule *model.ProximityRule, isAdmin bool) *httptest.ResponseRecorder {
	payload, _ := json.Marshal(proximityRule)
	req, _ := http.NewRequest(
		http.MethodPatch,
		"/recommended/rules/proximity",
		bytes.NewBuffer(payload),
	)
	userInfo := getUserInfo(isAdmin)
	req.Header.Set("user_info", userInfo)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	return recorder
}

func sendGetProximityRequest(router *gin.Engine, isAdmin bool) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(
		http.MethodGet,
		"/recommended/rules/proximity",
		nil,
	)

	userInfo := getUserInfo(isAdmin)
	req.Header.Set("user_info", userInfo)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	return recorder
}

func TestUpdateProximityRule(t *testing.T) {
	router, _ := setUpRouter(false, false)

	proximityRule := model.ProximityRule{
		Radius:  3,
		Enabled: true,
	}

	recorder := sendPatchProximityRequest(router, &proximityRule, true)

	assert.Equal(t, 201, recorder.Code)
}

func TestGetProximityRule(t *testing.T) {
	router, _ := setUpRouter(false, false)

	recorder := sendGetProximityRequest(router, true)

	assert.Equal(t, 200, recorder.Code)
}

func TestUpdateProximityRuleError(t *testing.T) {
	router, _ := setUpErrorRouter()

	proximityRule := model.ProximityRule{
		Radius:  3,
		Enabled: true,
	}
	recorder := sendPatchProximityRequest(router, &proximityRule, true)

	assert.Equal(t, 500, recorder.Code)
}

func TestGetProximityRuleError(t *testing.T) {
	router, _ := setUpErrorRouter()

	recorder := sendGetProximityRequest(router, true)

	assert.Equal(t, 500, recorder.Code)
}

func TestUpdateProximityRuleReturns401IfNotAdmin(t *testing.T) {
	router, _ := setUpRouter(false, false)

	proximityRule := model.ProximityRule{
		Radius:  3,
		Enabled: true,
	}

	recorder := sendPatchProximityRequest(router, &proximityRule, false)

	assert.Equal(t, 401, recorder.Code)
}

func TestGetProximityRuleReturns401IfNotAdmin(t *testing.T) {
	router, _ := setUpRouter(false, false)

	recorder := sendGetProximityRequest(router, false)

	assert.Equal(t, 401, recorder.Code)
}
