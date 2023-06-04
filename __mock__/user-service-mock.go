package __mock__

import "github.com/T2-1c2023/RecommendationService/app/model"

type UserServiceMock struct{}

func NewUserServiceMock() UserServiceMock {
	return UserServiceMock{}
}

func (service *UserServiceMock) GetUserById(id int, userInfo string) (model.User, error) {
	interests := []model.Interest{
		{Id: 1, Description: "Running"},
	}
	user := model.User{
		Id:            1,
		Fullname:      "John Doe",
		Mail:          "user@example.com",
		PhoneNumber:   "1234567890",
		IsAthlete:     true,
		IsTrainer:     true,
		IsVerified:    true,
		IsAdmin:       false,
		PhotoId:       "",
		ExpoPushToken: "",
		Latitude:      "0",
		Longitude:     "0",
		Interests:     interests,
	}
	return user, nil
}

func (service *UserServiceMock) GetAllTrainersInProximity(radius uint,
	userInfo string) ([]model.User, error) {
	user := model.User{
		Id:            1,
		Fullname:      "John Doe",
		Mail:          "user@example.com",
		PhoneNumber:   "1234567890",
		IsAthlete:     false,
		IsTrainer:     true,
		IsVerified:    true,
		IsAdmin:       false,
		PhotoId:       "",
		ExpoPushToken: "",
		Latitude:      "0",
		Longitude:     "0",
		Interests:     []model.Interest{},
	}
	return []model.User{user}, nil
}
