package validation

import (
	"net/http"

	"github.com/T2-1c2023/RecommendationService/app/model"
	"github.com/gin-gonic/gin"
)

func AdminValidator(c *gin.Context) {
	user, err := model.NewUserFromUserInfo(c.GetHeader("user_info"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if !user.IsAdmin {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": "Unauthorized"})
		return
	}
	c.Next()
}
