package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

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

func (service *TrainingService) GetTrainingsFromTrainerId(id int,
	userInfo string, queryParams map[string]string) ([]model.Training, error) {
	trainings := []model.Training{}
	req_url := fmt.Sprintf("%s/trainers/%d/trainings",
		os.Getenv("TRAININGS_MICROSERVICE_URL"),
		id)
	req, err := http.NewRequest("GET", req_url, nil)
	if err != nil {
		return trainings, fmt.Errorf("error creating a new request")
	}

	service.addQueryParams(req, queryParams)

	req.Header.Set("user_info", userInfo)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return trainings, fmt.Errorf("can't reach trainings microservice")
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&trainings)
	if err != nil {
		return trainings, fmt.Errorf("internal server error: " + err.Error())
	}
	return trainings, nil
}

func (service *TrainingService) GetAllTrainings(userInfo string,
	queryParams map[string]string) ([]model.Training, error) {
	trainings := []model.Training{}
	url := fmt.Sprintf("%s/trainings", os.Getenv("TRAININGS_MICROSERVICE_URL"))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return trainings, fmt.Errorf("error creating a new request")
	}

	service.addQueryParams(req, queryParams)

	req.Header.Set("user_info", userInfo)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return trainings, fmt.Errorf("can't reach users microservice")
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&trainings)
	if err != nil {
		return trainings, fmt.Errorf("internal server error: " + err.Error())
	}
	return trainings, nil
}
