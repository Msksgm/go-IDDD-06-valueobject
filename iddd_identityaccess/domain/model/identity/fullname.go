package identity

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
	if aFirstName == "" {
		return nil, fmt.Errorf("First name is required.")
	}
	if len(aFirstName) < 1 || 50 < len(aFirstName) {
		return nil, fmt.Errorf("First name must be 50 characters or less.")
	}
	if !regexp.MustCompile(`^[A-Z][a-z]*`).MatchString(aFirstName) {
		return nil, fmt.Errorf("First name must be at least one character in length, starting with a capital letter.")
	}
	fullName.firstName = aFirstName

	// set lastName
	if aLastName == "" {
		return nil, fmt.Errorf("Last name is required.")
	}
	if len(aLastName) < 1 || 50 < len(aLastName) {
		return nil, fmt.Errorf("Last name must be 50 characters or less.")
	}
	if !regexp.MustCompile(`^[a-zA-Z'][ a-zA-Z'-]*[a-zA-Z']?`).MatchString(aLastName) {
		return nil, fmt.Errorf("Last name must be at least one character in length.")
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
