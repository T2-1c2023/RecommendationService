// @title		Recommendation Microservice API
// @version         1.0
// @description     Recommendation microservice for FiuFit's backend.

package main

import (
	"os"

	"github.com/T2-1c2023/RecommendationService/app/controller"
	"github.com/T2-1c2023/RecommendationService/app/persistence"
	routes "github.com/T2-1c2023/RecommendationService/app/routes"
	"github.com/T2-1c2023/RecommendationService/app/services"
	"github.com/T2-1c2023/RecommendationService/app/utilities"
	config "github.com/T2-1c2023/RecommendationService/config"
	_ "github.com/T2-1c2023/RecommendationService/docs"
)

func main() {
	logDir := "./log"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		err := os.Mkdir(logDir, 0755)
		if err != nil {
			panic(err)
		}
	}
	logFile, err := os.OpenFile(logDir+"/app.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	logger := utilities.NewLogger(os.Getenv("LOG_LEVEL"), logFile)
	defer logger.CloseFile()
	logger.LogDebug("Microservice starting...")

	client, err := config.SetUpDB()
	if err != nil {
		panic(err)
	}
	logger.LogDebug("Database ready")

	repo := persistence.RulesRepository{
		Collection: client.Database("recDB").Collection("rules"),
	}
	userService := services.UserService{}
	trainingService := services.TrainingService{}

	proximityController := controller.ProximityRuleController{
		Repo:   &repo,
		Logger: &logger,
	}
	interestsController := controller.InterestsRuleController{
		Repo:   &repo,
		Logger: &logger,
	}
	recommendationController := controller.RecommendationController{
		Repo:            &repo,
		UserService:     &userService,
		TrainingService: &trainingService,
		Logger:          &logger,
	}
	statusController := controller.StatusController{
		Logger: &logger,
	}

	router := routes.SetupRouter(
		&proximityController,
		&interestsController,
		&recommendationController,
		&statusController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3003" // Default port
	}

	logger.LogInfo("App started on " + utilities.GetCurrentDate())
	logger.LogInfo("Recommendation microservice running at port " + port)
	router.Run(":" + port)
}
