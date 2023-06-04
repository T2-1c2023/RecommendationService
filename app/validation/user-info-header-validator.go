package validation

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserInfoHeaderValidator(c *gin.Context) {
	headerValue := c.GetHeader("user_info")
	if headerValue == "" {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"message": "Missing user info header"},
		)
	}
	c.Next()
}
