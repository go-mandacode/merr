package merr

import "google.golang.org/grpc/codes"

// error map
var grpcErrorMap = map[ErrCode]codes.Code{
	ErrUnknown:            codes.Unknown,
	ErrInternal:           codes.Internal,
	ErrNotFound:           codes.NotFound,
	ErrInvalidInput:       codes.InvalidArgument,
	ErrPermissionDenied:   codes.PermissionDenied,
	ErrTimeout:            codes.DeadlineExceeded,
	ErrConflict:           codes.Aborted,
	ErrUnauthorized:       codes.Unauthenticated,
	ErrBadRequest:         codes.InvalidArgument,
	ErrServiceUnavailable: codes.Unavailable,
	ErrTooManyRequests:    codes.ResourceExhausted,
	ErrNotImplemented:     codes.Unimplemented,
	ErrForbidden:          codes.PermissionDenied,
	ErrLoopDetected:       codes.Aborted,
}

func (e ErrCode) ToGRPCCode() codes.Code {
	if code, exists := grpcErrorMap[e]; exists {
		return code
	}
	return codes.Unknown // Default to Unknown if the error code is not mapped
}
