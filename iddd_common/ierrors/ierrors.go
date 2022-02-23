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

type ArgumentLengthErrorArguments struct {
	String  string
	Minimum int
	Maximum int
	Message string
}

type ArgumentLengthError struct {
	Arguments ArgumentLengthErrorArguments
}

func (ArgumentLengthErrorArguments *ArgumentLengthError) GetArguments() ArgumentLengthErrorArguments {
	return ArgumentLengthErrorArguments.Arguments
}

func (ArgumentLengthErrorArguments *ArgumentLengthError) GetError() error {
	args := ArgumentLengthErrorArguments.Arguments
	length := len(args.String)
	if length < args.Minimum || length > args.Maximum {
		return ArgumentLengthErrorArguments
	}
	return nil
}

func (ArgumentLengthErrorArguments *ArgumentLengthError) Error() string {
	return ArgumentLengthErrorArguments.Arguments.Message
}

type ArgumentNotEmptyErrorArguments struct {
	String  string
	Message string
}

type ArgumentNotEmptyError struct {
	Arguments ArgumentNotEmptyErrorArguments
}

func (ArgumentNotEmptyErrorArguments *ArgumentNotEmptyError) GetArguments() ArgumentNotEmptyErrorArguments {
	return ArgumentNotEmptyErrorArguments.Arguments
}

func (ArgumentNotEmptyErrorArguments *ArgumentNotEmptyError) GetError() error {
	args := ArgumentNotEmptyErrorArguments.Arguments
	if strings.TrimSpace(args.String) == "" {
		return ArgumentNotEmptyErrorArguments
	}
	return nil
}

func (ArgumentNotEmptyErrorArguments *ArgumentNotEmptyError) Error() string {
	return ArgumentNotEmptyErrorArguments.Arguments.Message
}
