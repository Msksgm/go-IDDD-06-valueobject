package ierrors

import (
	"fmt"
	"strings"
)

func Wrap(errp *error, format string, args ...interface{}) {
	if *errp != nil {
		*errp = fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), *errp)
	}
}

func NewArgumentLengthError(aString string, aMinimum int, aMaximum int, aMessage string) *ArgumentLengthError {
	arguments := ArgumentLengthErrorArguments{String: aString, Minimum: aMinimum, Maximum: aMaximum, Message: aMessage}
	return &ArgumentLengthError{Arguments: arguments}
}

type ArgumentLengthErrorArguments struct {
	String  string
	Minimum int
	Maximum int
	Message string
}

type ArgumentLengthError struct {
	Arguments ArgumentLengthErrorArguments
}

func (ArgumentLengthError *ArgumentLengthError) GetArguments() ArgumentLengthErrorArguments {
	return ArgumentLengthError.Arguments
}

func (ArgumentLengthError *ArgumentLengthError) GetError() error {
	args := ArgumentLengthError.Arguments
	length := len(args.String)
	if length < args.Minimum || length > args.Maximum {
		return ArgumentLengthError
	}
	return nil
}

func (ArgumentLengthError *ArgumentLengthError) Error() string {
	return ArgumentLengthError.Arguments.Message
}

type ArgumentNotEmptyErrorArguments struct {
	String  string
	Message string
}

type ArgumentNotEmptyError struct {
	Arguments ArgumentNotEmptyErrorArguments
}

func (ArgumentNotEmptyError *ArgumentNotEmptyError) GetArguments() ArgumentNotEmptyErrorArguments {
	return ArgumentNotEmptyError.Arguments
}

func (ArgumentNotEmptyError *ArgumentNotEmptyError) GetError() error {
	args := ArgumentNotEmptyError.Arguments
	if strings.TrimSpace(args.String) == "" {
		return ArgumentNotEmptyError
	}
	return nil
}

func (ArgumentNotEmptyError *ArgumentNotEmptyError) Error() string {
	return ArgumentNotEmptyError.Arguments.Message
}
