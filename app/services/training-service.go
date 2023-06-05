package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"

	model "github.com/T2-1c2023/RecommendationService/app/model"
)

type ITrainingService interface {
	GetTrainingsFromTrainerId(id int, userInfo string,
		queryParams map[string]string) ([]model.Training, error)
	GetAllTrainings(userInfo string,
		queryParams map[string]string) ([]model.Training, error)
}

type TrainingService struct{}

func (service *TrainingService) addQueryParams(req *http.Request,
	queryParams map[string]string) {
	if len(queryParams) > 0 {
		params := url.Values{}
		for key, value := range queryParams {
			params.Add(key, value)
		}
		req.URL.RawQuery = params.Encode()
	}
}

func (service *TrainingService) sendRequest(userInfo string,
	requestURL string, queryParams map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating a new request")
	}

	service.addQueryParams(req, queryParams)

	req.Header.Set("user_info", userInfo)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't reach trainings microservice")
	}
	return response, nil
}

func (service *TrainingService) getErrorMessageFromResponse(response *http.Response) error {
	var errorResponse model.UserServiceErrorResponse
	json.NewDecoder(response.Body).Decode(&errorResponse)
	message := "trainings microservice returned "
	message += strconv.Itoa(response.StatusCode) + ": "
	message += errorResponse.Message
	return fmt.Errorf(message)
}

func (service *TrainingService) GetTrainingsFromTrainerId(id int,
	userInfo string, queryParams map[string]string) ([]model.Training, error) {
	trainings := []model.Training{}
	requestURL := fmt.Sprintf("%s/trainers/%d/trainings",
		os.Getenv("TRAININGS_MICROSERVICE_URL"),
		id)

	response, err := service.sendRequest(userInfo, requestURL, queryParams)
	if err != nil {
		return trainings, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return trainings, service.getErrorMessageFromResponse(response)
	}

	err = json.NewDecoder(response.Body).Decode(&trainings)
	if err != nil {
		return trainings, fmt.Errorf("internal server error: " + err.Error())
	}
	return trainings, nil
}

func (service *TrainingService) GetAllTrainings(userInfo string,
	queryParams map[string]string) ([]model.Training, error) {
	trainings := []model.Training{}
	requestURL := fmt.Sprintf("%s/trainings", os.Getenv("TRAININGS_MICROSERVICE_URL"))

	response, err := service.sendRequest(userInfo, requestURL, queryParams)
	if err != nil {
		return trainings, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return trainings, service.getErrorMessageFromResponse(response)
	}

	err = json.NewDecoder(response.Body).Decode(&trainings)
	if err != nil {
		return trainings, fmt.Errorf("internal server error: " + err.Error())
	}
	return trainings, nil
}
