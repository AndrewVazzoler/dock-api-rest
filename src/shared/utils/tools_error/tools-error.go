package tools_error

import (
	"errors"
	"net/http"

	"github.com/jackc/pgconn"
)

type ToolsErrorsInterface interface {
	Transform(err error) *ErrorsDTO
}

type ErrorDBDTO struct {
	Code       string
	Message    string
	StatusCode int
	Cause      ErrorCauseDTO
}

type ErrorCauseDTO struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorsDTO struct {
	Code       string      `json:"code"`
	Message    string      `json:"message"`
	Cause      interface{} `json:"cause"`
	StatusCode int         `json:"status_code"`
}

type ToolsErrors struct {
}

func NewToolsErrors() *ToolsErrors {
	return &ToolsErrors{}
}

func (t *ToolsErrors) Transform(err error) *ErrorsDTO {
	var code = ErrorCodeInternal
	var message = "internal server error"
	var statusCode = http.StatusInternalServerError
	var cause interface{} = []ErrorCauseDTO{
		{
			Code:    code,
			Message: err.Error(),
		},
	}
	if dbErr := t.DBError(err); dbErr != nil {
		return &ErrorsDTO{
			Code:       dbErr.Code,
			Message:    dbErr.Message,
			StatusCode: dbErr.StatusCode,
			Cause: []ErrorCauseDTO{
				dbErr.Cause,
			},
		}
	}

	switch err.Error() {
	case ErrNotFound:
		code = DBErrorCodeNotFound
		message = "record not found"
		statusCode = http.StatusNotFound
	}

	return &ErrorsDTO{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
		Cause:      cause,
	}
}

func (t *ToolsErrors) DBError(err error) *ErrorDBDTO {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		var codeErro string = DBErrorCode
		var msg = "Database error"
		var status = http.StatusInternalServerError

		if pgErr.Code == "23505" {
			codeErro = DBErrorCodeDuplicate
			msg = pgErr.Detail
			status = http.StatusConflict
		}

		return &ErrorDBDTO{
			Code:       codeErro,
			Message:    msg,
			StatusCode: status,
			Cause: ErrorCauseDTO{
				Code:    pgErr.Code,
				Message: pgErr.Message,
			},
		}
	}
	return nil
}

func Transform(err error) *ErrorsDTO {
	return NewToolsErrors().Transform(err)
}
