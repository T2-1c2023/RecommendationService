package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	model "github.com/T2-1c2023/RecommendationService/app/model"
)

type UserServiceErrorResponse struct {
	Message string `json:"message" binding:"required"`
}

type IUserService interface {
	GetUserById(id int, userInfo string) (model.User, error)
	GetAllTrainersInProximity(radius uint,
		userInfo string) ([]model.User, error)
}

type UserService struct{}

func (service *UserService) GetUserById(id int, userInfo string) (model.User, error) {
	var user model.User
	url := os.Getenv("USERS_MICROSERVICE_URL") + "/users/" + strconv.Itoa(id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return user, fmt.Errorf("error creating a new request")
	}
	req.Header.Set("user_info", userInfo)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return user, fmt.Errorf("can't reach users microservice")
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return user, fmt.Errorf("internal server error: " + err.Error())
	}

	if response.StatusCode != http.StatusOK {
		var errorResponse UserServiceErrorResponse
		err = json.Unmarshal(bodyBytes, &errorResponse)
		if err != nil {
			return user, fmt.Errorf("internal server error: " + err.Error())
		}
		return user, fmt.Errorf("users microserviced returned " +
			strconv.Itoa(response.StatusCode) +
			": " +
			errorResponse.Message)
	}

	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		return user, fmt.Errorf("internal server error" + err.Error())
	}
	return user, nil
}

func (user *UserService) GetAllTrainersInProximity(radius uint,
	userInfo string) ([]model.User, error) {
	var trainers []model.User
	url := fmt.Sprintf("%s/users?is_trainer=true&radius=%d",
		os.Getenv("USERS_MICROSERVICE_URL"), radius)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return trainers, fmt.Errorf("error creating a new request")
	}
	req.Header.Set("user_info", userInfo)
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return trainers, fmt.Errorf("can't reach users microservice")
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(&trainers)
	if err != nil {
		return trainers, fmt.Errorf("internal server error: " + err.Error())
	}
	return trainers, nil
}
