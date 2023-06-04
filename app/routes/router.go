package routes

import (
	controller "github.com/T2-1c2023/RecommendationService/app/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(proximityController *controller.ProximityRuleController,
	interestsController *controller.InterestsRuleController) *gin.Engine {
	// Create a new Gin router
	router := gin.Default()

	router.GET("/", controller.GetStatus)

	router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/health", controller.GetHealth)

	router.PATCH("/rules/proximity", proximityController.ModifyProximityRule)
	router.GET("/rules/proximity", proximityController.GetProximityRule)

	router.PATCH("/rules/interests", interestsController.ModifyInterestsRule)
	router.GET("/rules/interests", interestsController.GetInterestsRule)

	return router
}
