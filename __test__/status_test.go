// __test__/app_test.go
package __test__

// func TestGetStatus(t *testing.T) {
// 	router, _ := setUpRouter(false, false)

// 	req, _ := http.NewRequest(http.MethodGet, "/", nil)
// 	recorder := httptest.NewRecorder()
// 	router.ServeHTTP(recorder, req)

// 	assert.Equal(t, 200, recorder.Code)
// }

// func TestGetHealth(t *testing.T) {
// 	router, _ := setUpRouter(false, false)

// 	req, _ := http.NewRequest(http.MethodGet, "/health", nil)
// 	recorder := httptest.NewRecorder()
// 	router.ServeHTTP(recorder, req)

// 	assert.Equal(t, 200, recorder.Code)
// }

// func TestBlockService(t *testing.T) {
// 	router, _ := setUpRouter(false, false)

// 	changeStatusInput := controller.ChangeStatusInput{
// 		Blocked: true,
// 	}
// 	payload, _ := json.Marshal(changeStatusInput)

// 	req, _ := http.NewRequest(http.MethodPost, "/status", bytes.NewBuffer(payload))
// 	userInfo := getUserInfo(true)
// 	req.Header.Set("user_info", userInfo)
// 	recorder := httptest.NewRecorder()
// 	router.ServeHTTP(recorder, req)

// 	assert.Equal(t, http.StatusOK, recorder.Code)
// }

// func TestGetBlockedStatus(t *testing.T) {
// 	router, _ := setUpRouter(false, false)

// 	req, _ := http.NewRequest(http.MethodGet, "/status", nil)
// 	recorder := httptest.NewRecorder()
// 	userInfo := getUserInfo(true)
// 	req.Header.Set("user_info", userInfo)
// 	router.ServeHTTP(recorder, req)

// 	finalChangeStatus := controller.ChangeStatusInput{
// 		Blocked: false,
// 	}

// 	var actualChangeStatus controller.ChangeStatusInput
// 	json.NewDecoder(recorder.Body).Decode(&actualChangeStatus)

// 	assert.Equal(t, http.StatusOK, recorder.Code)
// 	assert.Equal(t, finalChangeStatus, actualChangeStatus)
// }
