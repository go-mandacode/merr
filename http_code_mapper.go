package merr

import "net/http"

// error map
var httpErrorMap = map[ErrCode]int{
	ErrUnknown:            http.StatusInternalServerError,
	ErrNotFound:           http.StatusNotFound,
	ErrInvalidInput:       http.StatusBadRequest,
	ErrPermissionDenied:   http.StatusForbidden,
	ErrInternal:           http.StatusInternalServerError,
	ErrTimeout:            http.StatusGatewayTimeout,
	ErrConflict:           http.StatusConflict,
	ErrUnauthorized:       http.StatusUnauthorized,
	ErrBadRequest:         http.StatusBadRequest,
	ErrServiceUnavailable: http.StatusServiceUnavailable,
	ErrTooManyRequests:    http.StatusTooManyRequests,
	ErrNotImplemented:     http.StatusNotImplemented,
	ErrForbidden:          http.StatusForbidden,
	ErrLoopDetected:       http.StatusLoopDetected,
}

func (e ErrCode) ToHTTPStatus() int {
	if status, exists := httpErrorMap[e]; exists {
		return status
	}
	return http.StatusInternalServerError // Default to Internal Server Error if the error code is not mapped
}
