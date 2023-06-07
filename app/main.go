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
	config "github.com/T2-1c2023/RecommendationService/config"
	_ "github.com/T2-1c2023/RecommendationService/docs"
	// "github.com/joho/godotenv"
)

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	client, err := config.SetUpDB()
	if err != nil {
		panic(err)
	}

	repo := persistence.RulesRepository{
		Collection: client.Database("recDB").Collection("rules"),
	}
	userService := services.UserService{}
	trainingService := services.TrainingService{}

	proximityController := controller.ProximityRuleController{
		Repo: &repo,
	}
	interestsController := controller.InterestsRuleController{
		Repo: &repo,
	}
	recommendationController := controller.RecommendationController{
		Repo:            &repo,
		UserService:     &userService,
		TrainingService: &trainingService,
	}

	router := routes.SetupRouter(
		&proximityController,
		&interestsController,
		&recommendationController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3003" // Default port
	}

	router.Run(":" + port)
}
