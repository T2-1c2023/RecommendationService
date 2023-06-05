package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	model "github.com/T2-1c2023/RecommendationService/app/model"
)

type IUserService interface {
	GetUserById(id int, userInfo string) (model.User, error)
	GetAllTrainersInProximity(radius uint,
		userInfo string) ([]model.User, error)
}

type UserService struct{}

func (service *UserService) sendRequest(userInfo string,
	requestURL string) (*http.Response, error) {
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating a new request")
	}

	req.Header.Set("user_info", userInfo)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("can't reach users microservice")
	}
	return response, nil
}

func (service *UserService) getErrorMessageFromResponse(response *http.Response) error {
	var errorResponse model.UserServiceErrorResponse
	json.NewDecoder(response.Body).Decode(&errorResponse)
	message := "users microservice returned "
	message += strconv.Itoa(response.StatusCode) + ": "
	message += errorResponse.Message
	return fmt.Errorf(message)
}

func (service *UserService) GetUserById(id int, userInfo string) (model.User, error) {
	var user model.User
	requestURL := os.Getenv("USERS_MICROSERVICE_URL") + "/users/" + strconv.Itoa(id)
	response, err := service.sendRequest(userInfo, requestURL)
	if err != nil {
		return user, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return user, service.getErrorMessageFromResponse(response)
	}

	err = json.NewDecoder(response.Body).Decode(&user)
	if err != nil {
		return user, fmt.Errorf("internal server error: " + err.Error())
	}
	return user, nil
}

func (service *UserService) GetAllTrainersInProximity(radius uint,
	userInfo string) ([]model.User, error) {
	var trainers []model.User
	requestURL := fmt.Sprintf("%s/users?is_trainer=true&radius=%d",
		os.Getenv("USERS_MICROSERVICE_URL"), radius)
	response, err := service.sendRequest(userInfo, requestURL)
	if err != nil {
		return trainers, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return trainers, service.getErrorMessageFromResponse(response)
	}

	err = json.NewDecoder(response.Body).Decode(&trainers)
	if err != nil {
		return trainers, fmt.Errorf("internal server error: " + err.Error())
	}
	return trainers, nil
}
