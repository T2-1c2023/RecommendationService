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

func sendPatchInterestsRule(router *gin.Engine,
	interestsRule *model.InterestsRule, isAdmin bool) *httptest.ResponseRecorder {
	payload, _ := json.Marshal(interestsRule)
	req, _ := http.NewRequest(
		http.MethodPatch,
		"/recommended/rules/interests",
		bytes.NewBuffer(payload),
	)

	userInfo := getUserInfo(isAdmin)
	req.Header.Set("user_info", userInfo)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)
	return recorder
}

func sendGetInterestsRule(router *gin.Engine, isAdmin bool) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodGet, "/recommended/rules/interests", nil)

	userInfo := getUserInfo(isAdmin)
	req.Header.Set("user_info", userInfo)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	return recorder
}

func TestUpdateInterestsRule(t *testing.T) {
	router, _ := setUpRouter(false, false)

	interestsRule := model.InterestsRule{
		Enabled: true,
	}

	recorder := sendPatchInterestsRule(router, &interestsRule, true)

	assert.Equal(t, 201, recorder.Code)
}

func TestGetInterestsRule(t *testing.T) {
	router, _ := setUpRouter(false, false)

	recorder := sendGetInterestsRule(router, true)

	assert.Equal(t, 200, recorder.Code)
}

func TestUpdateInterestsRuleError(t *testing.T) {
	router, _ := setUpErrorRouter()

	interestsRule := model.InterestsRule{
		Enabled: true,
	}
	recorder := sendPatchInterestsRule(router, &interestsRule, true)

	assert.Equal(t, 500, recorder.Code)
}

func TestGetInterestsRuleError(t *testing.T) {
	router, _ := setUpErrorRouter()

	recorder := sendGetInterestsRule(router, true)

	assert.Equal(t, 500, recorder.Code)
}

func TestUpdateInterestsRuleReturns401IfNotAdmin(t *testing.T) {
	router, _ := setUpRouter(false, false)

	interestsRule := model.InterestsRule{
		Enabled: true,
	}

	recorder := sendPatchInterestsRule(router, &interestsRule, false)

	assert.Equal(t, 401, recorder.Code)
}

func TestGetInterestsRuleReturns401IfNotAdmin(t *testing.T) {
	router, _ := setUpRouter(false, false)

	recorder := sendGetInterestsRule(router, false)

	assert.Equal(t, 401, recorder.Code)
}
