package tools_dto

import "github.com/dranikpg/dto-mapper"

type ToolsDTO struct{}

func (t *ToolsDTO) Map(to, from interface{}) (err error) {
	err = dto.Map(to, from)
	return err
}

type ToolsDTOInterface interface {
	Map(to, from interface{}) (err error)
}

func NewToolsDTO() *ToolsDTO {
	return &ToolsDTO{}
}
