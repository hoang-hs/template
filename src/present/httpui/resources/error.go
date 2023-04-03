package resources

import (
	"base/src/common"
	"base/src/core/constant"
	"net/http"
)

type ErrorResponse struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	TraceID    string `json:"trace_id,omitempty"`
	Detail     string `json:"detail,omitempty"`
	Source     string `json:"source"`
	HTTPStatus int    `json:"http_status"`
}

func ConvertErrorToResponse(err *common.Error) *ErrorResponse {
	detail := ""
	if !isInternalError(err) || !constant.IsProdEnv() {
		detail = err.Detail
	}
	return &ErrorResponse{
		Code:       int(err.Code),
		Message:    err.Message,
		TraceID:    err.TraceID,
		Detail:     detail,
		Source:     string(err.Source),
		HTTPStatus: err.HTTPStatus,
	}
}

func isInternalError(err *common.Error) bool {
	return err.GetHttpStatus() >= http.StatusInternalServerError
}
