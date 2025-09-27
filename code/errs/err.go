package errs

import "errors"

var (
	ErrTooManyRequest = errors.New("too Many Request")
	ErrWrongCode      = errors.New("wrong Code")
	ErrSystemError    = errors.New("code system error")
)
