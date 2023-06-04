package __mock__

import "github.com/T2-1c2023/RecommendationService/app/model"

type TrainingServiceMock struct {
	Trainings []model.Training
}

func NewTrainingServiceMock() TrainingServiceMock {
	allTrainings := []model.Training{
		{
			Id:    1,
			Title: "Test training",
			Type:  "Running",
		},
		{
			Id:    2,
			Title: "Test training 2",
			Type:  "Caminata",
		},
		{
			Id:    3,
			Title: "Test training 3",
			Type:  "Running",
		},
	}
	return TrainingServiceMock{
		Trainings: allTrainings,
	}
}

func (service *TrainingServiceMock) GetTrainingsFromTrainerId(id int, userInfo string) ([]model.Training, error) {
	return service.Trainings[:2], nil
}

func (service *TrainingServiceMock) GetAllTrainings(userInfo string) ([]model.Training, error) {
	return service.Trainings, nil
}
