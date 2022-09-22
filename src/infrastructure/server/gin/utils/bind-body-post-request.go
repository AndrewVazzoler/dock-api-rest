package utils

import (
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/server/gin/validation"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func BindBodyFromPostRequest(c *gin.Context, obj interface{}) *validation.ErrorRoot {
	if err := c.ShouldBindBodyWith(&obj, binding.JSON); err != nil {
		v := validation.Translate(err)

		if len(v) > 0 {
			return &validation.ErrorRoot{Error: v}
		}
		return &validation.ErrorRoot{Error: err.Error()}
	}
	return nil
}
