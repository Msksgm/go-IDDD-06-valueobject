package ierrors

import (
	"fmt"
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
