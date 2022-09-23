package tools_http

import (
	"github.com/AndrewVazzoler/dock-api-rest/src/infrastructure/server/gin/validation"

	"net/http"

	"github.com/AndrewVazzoler/dock-api-rest/src/shared/utils/tools_error"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type ToolsHttpInterface interface {
	BadRequest(err error, obj interface{})
	ResponseError(err error)
	OK(obj interface{})
	DeleteOK()
	CreateOK(obj interface{})
	SetServerContext(context *gin.Context)
	BindBodyFromPostRequest(obj interface{}) *tools_error.ErrorsDTO
}

type ToolsHttp struct {
	context *gin.Context
}

func NewToolsHttp() *ToolsHttp {
	return &ToolsHttp{}
}

func (t *ToolsHttp) SetServerContext(context *gin.Context) {
	t.context = context
}

func (t *ToolsHttp) BindBodyFromPostRequest(obj interface{}) *tools_error.ErrorsDTO {
	if err := t.context.ShouldBindBodyWith(&obj, binding.JSON); err != nil {
		v := validation.Translate(err)

		var cause interface{} = []tools_error.ErrorCauseDTO{
			{
				Code:    tools_error.ErrorCodeUnmarshal,
				Message: err.Error(),
			},
		}
		if len(v) > 0 {
			cause = v
		}
		return &tools_error.ErrorsDTO{
			Code:       tools_error.ErrorCodeValidations,
			Message:    tools_error.ErrorValidations,
			StatusCode: http.StatusBadRequest,
			Cause:      cause,
		}
	}
	return nil
}

func (t *ToolsHttp) OK(obj interface{}) {
	t.context.JSON(http.StatusCreated, obj)
}

func (t *ToolsHttp) DeleteOK() {
	t.context.AbortWithStatus(http.StatusNoContent)
}

func (t *ToolsHttp) BadRequest(err error, obj interface{}) {
	t.context.JSON(http.StatusBadRequest, &tools_error.ErrorsDTO{
		Code:       tools_error.ErrorCodeBadRequest,
		Message:    err.Error(),
		Cause:      obj,
		StatusCode: http.StatusBadRequest,
	})
}

func (t *ToolsHttp) ResponseError(err error) {
	e := tools_error.Transform(err)
	t.context.JSON(e.StatusCode, e)
}

func (t *ToolsHttp) CreateOK(obj interface{}) {
	t.context.JSON(http.StatusCreated, obj)
}
