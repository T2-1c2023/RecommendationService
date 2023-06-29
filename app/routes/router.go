package routes

import (
	controller "github.com/T2-1c2023/RecommendationService/app/controller"
	"github.com/T2-1c2023/RecommendationService/app/validation"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(proximityController *controller.ProximityRuleController,
	interestsController *controller.InterestsRuleController,
	recommendationController *controller.RecommendationController,
	statusController *controller.StatusController) *gin.Engine {
	// Create a new Gin router
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	router.Use(cors.New(config))

	router.GET("/", statusController.GetStatus)

	router.GET("/status",
		validation.UserInfoHeaderValidator,
		validation.AdminValidator,
		statusController.GetServiceStatus)
	router.POST("/status",
		validation.UserInfoHeaderValidator,
		validation.AdminValidator,
		statusController.ChangeServiceStatus)

	router.GET("/api-docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/logs", statusController.GetLogs)

	router.Use(statusController.ValidateBlockedStatus)

	router.GET("/health", statusController.GetHealth)

	router.PATCH("/recommended/rules/proximity",
		validation.UserInfoHeaderValidator,
		validation.AdminValidator,
		proximityController.ModifyProximityRule)
	router.GET("/recommended/rules/proximity",
		validation.AdminValidator,
		proximityController.GetProximityRule)

	router.PATCH("/recommended/rules/interests",
		validation.UserInfoHeaderValidator,
		validation.AdminValidator,
		interestsController.ModifyInterestsRule)
	router.GET("/recommended/rules/interests",
		validation.UserInfoHeaderValidator,
		validation.AdminValidator,
		interestsController.GetInterestsRule)

	router.GET("/recommended",
		validation.UserInfoHeaderValidator,
		recommendationController.GetRecommendations)

	return router
}
