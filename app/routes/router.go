package routes

import (
	controller "github.com/T2-1c2023/RecommendationService/app/controller"
	"github.com/T2-1c2023/RecommendationService/app/validation"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(proximityController *controller.ProximityRuleController,
	interestsController *controller.InterestsRuleController,
	recommendationController *controller.RecommendationController) *gin.Engine {
	// Create a new Gin router
	router := gin.Default()

	router.GET("/", controller.GetStatus)

	router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/health", controller.GetHealth)

	router.PATCH("/rules/proximity",
		validation.UserInfoHeaderValidator,
		validation.AdminValidator,
		proximityController.ModifyProximityRule)
	router.GET("/rules/proximity",
		validation.AdminValidator,
		proximityController.GetProximityRule)

	router.PATCH("/rules/interests",
		validation.UserInfoHeaderValidator,
		validation.AdminValidator,
		interestsController.ModifyInterestsRule)
	router.GET("/rules/interests",
		validation.UserInfoHeaderValidator,
		validation.AdminValidator,
		interestsController.GetInterestsRule)

	router.GET("/recommended",
		validation.UserInfoHeaderValidator,
		recommendationController.GetRecommendations)

	return router
}
