package fullname

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
)

type FullName struct {
	firstName string
	lastName  string
}

func NewFullName(aFirstName string, aLastName string) (_ *FullName, err error) {
	defer ierrors.Wrap(&err, "fullname.NewFullName(%s, %s)", aFirstName, aLastName)
	fullName := new(FullName)

	// set firstName
	if err := ierrors.NewArgumentNotEmptyError(aFirstName, "First name is required.").GetError(); err != nil {
		return nil, err
	}
	if err := ierrors.NewArgumentLengthError(aFirstName, 1, 50, "First name must be 50 characters or less.").GetError(); err != nil {
		return nil, err
	}
	if err := ierrors.NewArgumentTrueErrorArguments(regexp.MustCompile(`^[A-Z][a-z]*`).MatchString(aFirstName), "First name must be at least one character in length, starting with a capital letter.").GetError(); err != nil {
		return nil, err
	}
	fullName.firstName = aFirstName

	// set lastName
	if err := ierrors.NewArgumentNotEmptyError(aLastName, "Last name is required.").GetError(); err != nil {
		return nil, err
	}
	if err := ierrors.NewArgumentLengthError(aLastName, 1, 50, "Last name must be 50 characters or less.").GetError(); err != nil {
		return nil, err
	}
	if err := ierrors.NewArgumentTrueErrorArguments(regexp.MustCompile(`^[a-zA-Z'][ a-zA-Z'-]*[a-zA-Z']?`).MatchString(aLastName), "Last name must be at least one character in length.").GetError(); err != nil {
		return nil, err
	}
	fullName.lastName = aLastName

	return fullName, nil
}

// TODO CopyFullName as shallow copy

func (fullName *FullName) AsFormattedName() string {
	return fmt.Sprintf("%s %s", fullName.firstName, fullName.lastName)
}

func (fullName *FullName) FirstName() string {
	return fullName.firstName
}

func (fullName *FullName) LastName() string {
	return fullName.lastName
}

func (fullName *FullName) WithChangedFirstName(aFirstName string) (_ *FullName, err error) {
	defer ierrors.Wrap(&err, "fullname.WithChangedFirstName(%s)", aFirstName)
	fullName, err = NewFullName(aFirstName, fullName.lastName)
	if err != nil {
		return nil, err
	}
	return fullName, nil
}

func (fullName *FullName) WithChangedLastName(aLastName string) (_ *FullName, err error) {
	defer ierrors.Wrap(&err, "fullname.WithChangedLastName(%s)", aLastName)
	fullName, err = NewFullName(fullName.firstName, aLastName)
	if err != nil {
		return nil, err
	}
	return fullName, nil
}

func (fullName *FullName) Equal(otherFullName *FullName) bool {
	isFirstNameEqual := reflect.DeepEqual(fullName.firstName, otherFullName.firstName)
	isLastNameEqual := reflect.DeepEqual(fullName.lastName, otherFullName.lastName)

	return isFirstNameEqual && isLastNameEqual
}

func (fullName *FullName) String() string {
	return fmt.Sprintf("FullName [firstName=" + fullName.firstName + ", lastName=" + fullName.lastName + "]")
}
