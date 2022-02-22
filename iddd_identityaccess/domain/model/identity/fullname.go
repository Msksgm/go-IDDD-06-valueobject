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

func NewFullName(firstName string, lastName string) (_ *FullName, err error) {
	defer ierrors.Wrap(&err, "fullname.NewFullName(%s, %s)", firstName, lastName)
	fullName := new(FullName)

	// set firstName
	if firstName == "" {
		return nil, fmt.Errorf("First name is required.")
	}
	if len(firstName) < 1 || 50 < len(firstName) {
		return nil, fmt.Errorf("First name must be 50 characters or less.")
	}
	if !regexp.MustCompile(`^[A-Z][a-z]*`).MatchString(firstName) {
		return nil, fmt.Errorf("First name must be at least one character in length, starting with a capital letter.")
	}
	fullName.firstName = firstName

	// set lastName
	if lastName == "" {
		return nil, fmt.Errorf("Last name is required.")
	}
	if len(lastName) < 1 || 50 < len(lastName) {
		return nil, fmt.Errorf("Last name must be 50 characters or less.")
	}
	if !regexp.MustCompile(`^[a-zA-Z'][ a-zA-Z'-]*[a-zA-Z']?`).MatchString(firstName) {
		return nil, fmt.Errorf("Last name must be at least one character in length.")
	}
	fullName.lastName = lastName

	return fullName, nil
}

func (fullName *FullName) Equal(otherFullName *FullName) bool {
	isFirstNameEqual := reflect.DeepEqual(fullName.firstName, otherFullName.firstName)
	isLastNameEqual := reflect.DeepEqual(fullName.lastName, otherFullName.lastName)

	return isFirstNameEqual && isLastNameEqual
}
