package errors

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

var (
	// ErrTetapTenangTetapSemangat custom error on unexpected error
	ErrTetapTenangTetapSemangat = CustomError{
		Mask:     "Tetap Tenang Tetap Semangat",
		Message:  "Tetap Tenang Tetap Semangat",
		HTTPCode: http.StatusInternalServerError,
	}

	ErrBadRequest = CustomError{
		Mask:     "Bad Request",
		Message:  "Bad Request",
		HTTPCode: http.StatusBadRequest,
	}

	ErrUnauthorized = CustomError{
		Mask:     "Unauthorized",
		Message:  "Unauthorized",
		HTTPCode: http.StatusUnauthorized,
	}

	ErrForbidden = CustomError{
		Mask:     "Forbidden",
		Message:  "Forbidden",
		HTTPCode: http.StatusForbidden,
	}

	ErrNotFound = CustomError{
		Mask:     "Record not exist",
		Message:  "Record not exist",
		HTTPCode: http.StatusNotFound,
	}

	ErrUnprocessableEntity = CustomError{
		Mask:     "Unprocessable Entity",
		Message:  "Unprocessable Entity",
		HTTPCode: http.StatusUnprocessableEntity,
	}

	ErrFailedAuthentication = CustomError{
		Mask:     "Invalid Credentials",
		Message:  "Invalid Credentials",
		HTTPCode: http.StatusUnauthorized,
	}

	ErrInternalServerError = CustomError{
		Mask:     "Internal Server Error",
		Message:  "Internal Server Error",
		HTTPCode: http.StatusInternalServerError,
	}
)

// CustomError holds data for customized error
type CustomError struct {
	ErrWithStack error       `json:"-"`
	Mask         interface{} `json:"-"`
	Message      interface{} `json:"message"`
	HTTPCode     int         `json:"code"`
	Report       bool        `json:"-"`
}

// Error is a function to convert error to string.
// It exists to satisfy error interface
func (c CustomError) Error() string {
	return fmt.Sprint(c.Message)
}

// ErrorMask is a function to return error mask.
func (c CustomError) ErrorMask() string {
	return fmt.Sprint(c.Mask)
}

// New, Errorf, Wrap, and Wrapf record a stack trace at the point they are invoked.
// StackTracer can be used to retrieve this information
type StackTracer interface {
	StackTrace() errors.StackTrace
}

// New returns an error with the supplied message.
// New also records the stack trace at the point it was called.
func New(message string) error {
	return errors.New(message)
}

func CustomWrap(err error) error {
	customErr, ok := err.(CustomError)
	if ok && customErr.ErrWithStack == nil {
		return CustomError{
			ErrWithStack: Wrap(err, customErr.ErrorMask()),
			HTTPCode:     customErr.HTTPCode,
			Mask:         customErr.ErrorMask(),
			Message:      customErr.Error(),
		}
	}
	return err
}

// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
// Errorf also records the stack trace at the point it was called.
func Errorf(format string, args ...interface{}) error {
	return errors.Errorf(format, args)
}

// WithStack annotates err with a stack trace at the point WithStack was called.
// If err is nil, WithStack returns nil.
func WithStack(err error) error {
	return errors.WithStack(err)
}

// Wrap returns an error annotating err with a stack trace
// at the point Wrap is called, and the supplied message.
// If err is nil, Wrap returns nil.
func Wrap(err error, message string) error {
	return errors.Wrap(err, message)
}

// Wrapf returns an error annotating err with a stack trace
// at the point Wrapf is called, and the format specifier.
// If err is nil, Wrapf returns nil.
func Wrapf(err error, format string, args ...interface{}) error {
	return errors.Wrapf(err, format, args)
}

// WithMessage annotates err with a new message.
// If err is nil, WithMessage returns nil.
func WithMessage(err error, message string) error {
	return errors.WithMessagef(err, message)
}

// WithMessagef annotates err with the format specifier.
// If err is nil, WithMessagef returns nil.
func WithMessagef(err error, format string, args ...interface{}) error {
	return errors.WithMessagef(err, format, args)
}

// Cause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//	type causer interface {
//	       Cause() error
//	}
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
func Cause(err error) error {
	return errors.Cause(err)
}
