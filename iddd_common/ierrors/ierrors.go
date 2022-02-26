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

func NewArgumentNotEmptyError(aString string, aMessage string) *ArgumentNotEmptyError {
	arguments := ArgumentNotEmptyErrorArguments{String: aString, Message: aMessage}
	return &ArgumentNotEmptyError{Arguments: arguments}
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

func NewArgumentTrueErrorArguments(aBool bool, aMessage string) *ArgumentTrueError {
	arguments := ArgumentTrueErrorArguments{Bool: aBool, Message: aMessage}
	return &ArgumentTrueError{Arguments: arguments}
}

type ArgumentTrueErrorArguments struct {
	Bool    bool
	Message string
}

type ArgumentTrueError struct {
	Arguments ArgumentTrueErrorArguments
}

func (ArgumentTrueError *ArgumentTrueError) GetArguments() ArgumentTrueErrorArguments {
	return ArgumentTrueError.Arguments
}

func (ArgumentTrueError *ArgumentTrueError) GetError() error {
	args := ArgumentTrueError.Arguments
	if !args.Bool {
		return ArgumentTrueError
	}
	return nil
}

func (ArgumentTrueError *ArgumentTrueError) Error() string {
	return ArgumentTrueError.Arguments.Message
}

func NewArgumentFalseError(aBool bool, aMessage string) *ArgumentFalseError {
	arguments := ArgumentFalseErrorArguments{isFalse: aBool, message: aMessage}
	return &ArgumentFalseError{arguments: arguments}
}

type ArgumentFalseErrorArguments struct {
	isFalse bool
	message string
}

type ArgumentFalseError struct {
	arguments ArgumentFalseErrorArguments
}

func (argumentFalseError *ArgumentFalseError) GetArguments() ArgumentFalseErrorArguments {
	return argumentFalseError.arguments
}

func (argumentFalseError *ArgumentFalseError) GetError() error {
	args := argumentFalseError.arguments
	if args.isFalse {
		return argumentFalseError
	}
	return nil
}

func (argumentFalseError *ArgumentFalseError) Error() string {
	return argumentFalseError.arguments.message
}
