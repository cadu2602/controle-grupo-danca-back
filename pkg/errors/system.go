package errors

import "strings"

const defaultSystemErrorCode = "internal"

type SystemError struct {
	BusinessError
	Reason error `json:"reason"`
}

func System(reason error, message, code string) SystemError {
	if code == "" {
		code = defaultSystemErrorCode
	}

	return SystemError{
		Reason:        reason,
		BusinessError: Business(message, code),
	}
}

func (e SystemError) WithErr(err error) SystemError {
	e.Reason = err

	return e
}

func (e SystemError) Error() string {
	var builder strings.Builder

	builder.WriteString(e.BusinessError.Error())

	if e.Reason == nil {
		return builder.String()
	}

	builder.WriteString(" (")
	builder.WriteString(e.Reason.Error())
	builder.WriteString(")")

	return builder.String()
}
