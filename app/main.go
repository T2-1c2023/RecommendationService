// @title		Recommendation Microservice API
// @version         1.0
// @description     Recommendation microservice for FiuFit's backend.

package main

import (
	"os"

	"github.com/T2-1c2023/RecommendationService/app/controller"
	"github.com/T2-1c2023/RecommendationService/app/persistence"
	routes "github.com/T2-1c2023/RecommendationService/app/routes"
	config "github.com/T2-1c2023/RecommendationService/config"
	_ "github.com/T2-1c2023/RecommendationService/docs"
)

func main() {
	client, err := config.SetUpDB()
	if err != nil {
		panic(err)
	}

	repo := persistence.RulesRepository{
		Collection: client.Database("recDB").Collection("rules"),
	}
	proximityController := controller.ProximityRuleController{
		Repo: &repo,
	}
	interestsController := controller.InterestsRuleController{
		Repo: &repo,
	}

	router := routes.SetupRouter(&proximityController, &interestsController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3003" // Default port
	}

	router.Run(":" + port)
}
