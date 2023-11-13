package error

import (
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type Diagnostic interface {
	String() string
	IsError() bool
}

type errorDiagnostic struct {
	err error
}

func (e *errorDiagnostic) String() string {
	return e.err.Error()
}

func (e *errorDiagnostic) IsError() bool {
	return true
}

type messageDiagnostic struct {
	message string
	isError bool
}

func (e *messageDiagnostic) String() string {
	return e.message
}

func (e *messageDiagnostic) IsError() bool {
	return e.isError
}

type Diagnostics []Diagnostic

func (d Diagnostics) AddError(err error) Diagnostics {

	if err == nil {
		return d
	}

	return append(d, &errorDiagnostic{err: err})

}

func (d Diagnostics) AddErrorWithMessage(err error, msg string) Diagnostics {

	if err == nil {
		return d
	}

	return append(d, &messageDiagnostic{message: msg, isError: true})

}

func (d Diagnostics) HasError() bool {

	for _, diagnostic := range d {
		if diagnostic.IsError() {
			return true
		}
	}

	return false

}

func (d Diagnostics) ToGrpcError(code codes.Code, requestIdHeader string) error {

	errors := []string{}

	for _, diagnostic := range d {
		if diagnostic.IsError() {
			errors = append(errors, diagnostic.String())
		}
	}

	if len(errors) == 0 {
		return nil
	}

	return grpc.Errorf(code, strings.Join(errors, "; "))

}
