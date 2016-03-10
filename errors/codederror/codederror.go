package codederror

import "strings"

// CodedError is a concrete implementation of codederror.Interface
type CodedError struct {
	Err  string `json:"error"`
	Code int    `json:"code"`
	Info string `json:"info,omitempty"`
}

// New initializes a new CodedError with optional info.
func New(err string, code int, info ...string) CodedError {
	i := strings.Join(info, "")
	return CodedError{
		Err:  err,
		Code: code,
		Info: i,
	}
}

// NewGeneric initializes a new CodedError for an existing error and sets the code to 900.
func NewGeneric(err error) CodedError {
	return CodedError{
		Err:  err.Error(),
		Code: CodeErrGeneric,
	}
}

// ErrorMessage satisfies codederror.Interface and the standard error by returning the error message.
func (c CodedError) Error() string {
	return c.Err
}

// ErrorCode satisfies codederror.Interface by returning the error code.
func (c CodedError) ErrorCode() int {
	return c.Code
}

// AdditionalInfo satisfies codederror.Interface by returning the additional info.
func (c CodedError) AdditionalInfo() string {
	return c.Info
}
