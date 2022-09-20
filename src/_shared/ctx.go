package shared

import (
	"github.com/AndrewVazzoler/dock-api-rest/src/_shared/utils/tools_dto"
	"github.com/AndrewVazzoler/dock-api-rest/src/_shared/utils/tools_error"
	"github.com/AndrewVazzoler/dock-api-rest/src/_shared/utils/tools_http"
)

type Ctx struct {
	ToolsDTO    tools_dto.ToolsDTOInterface
	ToolsErrors tools_error.ToolsErrorsInterface
	Http        tools_http.ToolsHttpInterface
}

func NewCtx() Ctx {
	return Ctx{
		ToolsDTO:    tools_dto.NewToolsDTO(),
		ToolsErrors: tools_error.NewToolsErrors(),
		Http:        tools_http.NewToolsHttp(),
	}
}
