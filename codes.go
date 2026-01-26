package merr

type ErrCode int

const (
	// ErrUnknown            ErrCode = "unknown"
	// ErrNotFound           ErrCode = "not_found"
	// ErrInternal           ErrCode = "internal_error"
	// ErrTimeout            ErrCode = "timeout"
	// ErrInvalidInput       ErrCode = "invalid_input"
	// ErrPermissionDenied   ErrCode = "permission_denied"
	// ErrForbidden          ErrCode = "forbidden"
	// ErrConflict           ErrCode = "conflict"
	// ErrUnauthorized       ErrCode = "unauthorized"
	// ErrBadRequest         ErrCode = "bad_request"
	// ErrServiceUnavailable ErrCode = "service_unavailable"
	// ErrTooManyRequests    ErrCode = "too_many_requests"
	// ErrNotImplemented     ErrCode = "not_implemented"
	// ErrLoopDetected       ErrCode = "loop_detected"
	ErrUnknown ErrCode = iota
	ErrInternal
	ErrNotFound
	ErrTimeout
	ErrInvalidInput
	ErrPermissionDenied
	ErrForbidden
	ErrConflict
	ErrUnauthorized
	ErrBadRequest
	ErrServiceUnavailable
	ErrTooManyRequests
	ErrNotImplemented
	ErrLoopDetected

	// errCodeEnd is a sentinel value for iteration. Must be the last constant.
	errCodeEnd
)

func CheckCode(err error, code ErrCode) bool {
	if err == nil {
		return false
	}
	if serr, ok := err.(PublicErr); ok {
		return serr.Code() == code
	}
	return false
}
