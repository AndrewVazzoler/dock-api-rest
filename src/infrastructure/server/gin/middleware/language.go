package middleware

import (
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/server/gin/validation"

	"github.com/gin-gonic/gin"
)

func DetectLanguage() gin.HandlerFunc {
	return func(c *gin.Context) {
		selectedLanguage := ""

		lang, b := c.Request.Header["Accept-Language"]

		if b && len(lang) > 0 {
			selectedLanguage = lang[0]
		}

		validation.Init(selectedLanguage)
		c.Next()
	}
}
