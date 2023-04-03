package common

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type CodeResponse int

const (
	//internal
	ErrorCodeBadRequest   CodeResponse = 400
	ErrorCodeUnauthorized CodeResponse = 401
	ErrorCodeSystemError  CodeResponse = 500
)

type Source string

const (
	SourceAPIService Source = "API_Service"
)

type Error struct {
	Code       CodeResponse `json:"code"`
	Message    string       `json:"message"`
	TraceID    string       `json:"trace_id,omitempty"`
	Detail     string       `json:"detail"`
	Source     Source       `json:"source"`
	HTTPStatus int          `json:"http_status"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("code:[%d], message:[%s], detail:[%s], source:[%s]", e.Code, e.Message, e.Detail, e.Source)
}

func (e *Error) GetHttpStatus() int {
	return e.HTTPStatus
}

func (e *Error) GetCode() CodeResponse {
	return e.Code
}

func (e *Error) GetMessage() string {
	return e.Message
}

func (e *Error) SetTraceId(traceId string) *Error {
	e.TraceID = fmt.Sprintf("%s:%d", traceId, time.Now().Unix())
	return e
}

func (e *Error) SetHTTPStatus(status int) *Error {
	e.HTTPStatus = status
	return e
}

func (e *Error) SetMessage(msg string) *Error {
	e.Message = msg
	return e
}

func (e *Error) SetDetail(detail string) *Error {
	e.Detail = detail
	return e
}

func (e *Error) GetDetail() string {
	return e.Detail
}

func (e *Error) SetSource(source Source) *Error {
	e.Source = source
	return e
}

func (e *Error) ToJSon() string {
	data, err := json.Marshal(e)
	if err != nil {
		//Todo fix this
		return "marshal error failed"
	}
	return string(data)
}

var (
	// Status 4xx ********

	ErrUnauthorized = func(ctx context.Context) *Error {
		traceId := GetTraceId(ctx)
		return &Error{
			Code:       ErrorCodeUnauthorized,
			Message:    DefaultUnauthorizedMessage,
			TraceID:    traceId,
			Source:     SourceAPIService,
			HTTPStatus: http.StatusUnauthorized,
		}
	}

	ErrBadRequest = func(ctx context.Context) *Error {
		traceId := GetTraceId(ctx)
		return &Error{
			Code:       ErrorCodeBadRequest,
			Message:    DefaultBadRequestMessage,
			TraceID:    traceId,
			HTTPStatus: http.StatusBadRequest,
			Source:     SourceAPIService,
		}
	}

	// Status 5xx *******

	ErrSystemError = func(ctx context.Context, detail string) *Error {
		traceId := GetTraceId(ctx)
		return &Error{
			Code:       ErrorCodeSystemError,
			Message:    DefaultServerErrorMessage,
			TraceID:    traceId,
			HTTPStatus: http.StatusInternalServerError,
			Source:     SourceAPIService,
			Detail:     detail,
		}
	}
)

const (
	DefaultServerErrorMessage  = "Something has gone wrong, please contact admin"
	DefaultBadRequestMessage   = "Invalid request"
	DefaultUnauthorizedMessage = "Token invalid"
)

func IsClientError(err *Error) bool {
	if err == nil {
		return false
	}
	if http.StatusBadRequest <= err.GetHttpStatus() && err.GetHttpStatus() < http.StatusInternalServerError {
		return true
	}
	return false
}

func IsInternalError(err *Error) bool {
	if err == nil {
		return false
	}
	if err.GetHttpStatus() >= http.StatusInternalServerError {
		return true
	}
	return false
}
