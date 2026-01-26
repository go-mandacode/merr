package merr

import (
	"testing"

	"google.golang.org/grpc/codes"
)

func TestAllErrCodesAreMappedToGRPC(t *testing.T) {
	for code := range errCodeEnd {
		if _, exists := grpcErrorMap[code]; !exists {
			t.Errorf("ErrCode %d is not mapped in grpcErrorMap", code)
		}
	}
}

func TestAllErrCodesAreMappedToHTTP(t *testing.T) {
	for code := range errCodeEnd {
		if _, exists := httpErrorMap[code]; !exists {
			t.Errorf("ErrCode %d is not mapped in httpErrorMap", code)
		}
	}
}

func TestGRPCMapDoesNotContainUnknownCodes(t *testing.T) {
	for code := range grpcErrorMap {
		if code < 0 || code >= errCodeEnd {
			t.Errorf("grpcErrorMap contains unknown ErrCode: %d", code)
		}
	}
}

func TestHTTPMapDoesNotContainUnknownCodes(t *testing.T) {
	for code := range httpErrorMap {
		if code < 0 || code >= errCodeEnd {
			t.Errorf("httpErrorMap contains unknown ErrCode: %d", code)
		}
	}
}

func TestToGRPCCodeReturnsValidCode(t *testing.T) {
	for code := ErrCode(0); code < errCodeEnd; code++ {
		// Ensure the method doesn't panic and returns a value
		_ = code.ToGRPCCode()
	}
}

func TestToHTTPStatusReturnsValidCode(t *testing.T) {
	for code := ErrCode(0); code < errCodeEnd; code++ {
		httpStatus := code.ToHTTPStatus()
		if httpStatus < 100 || httpStatus > 599 {
			t.Errorf("ErrCode %d returned invalid HTTP status: %d", code, httpStatus)
		}
	}
}

func TestUnmappedErrCodeFallsBackToDefault(t *testing.T) {
	unmappedCode := ErrCode(9999)

	grpcCode := unmappedCode.ToGRPCCode()
	if grpcCode != codes.Unknown {
		t.Errorf("Unmapped ErrCode should return codes.Unknown for gRPC, got: %v", grpcCode)
	}

	httpStatus := unmappedCode.ToHTTPStatus()
	if httpStatus != 500 {
		t.Errorf("Unmapped ErrCode should return 500 for HTTP, got: %d", httpStatus)
	}
}
