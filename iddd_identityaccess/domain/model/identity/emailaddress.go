package identity

import (
	"fmt"
	"regexp"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
)

type EmailAddress struct {
	address string
}

func NewEmailAddress(address string) (_ *EmailAddress, err error) {
	defer ierrors.Wrap(&err, "emailaddress.NewEmailAddress(%s)", address)
	emailAddress := new(EmailAddress)
	if address == "" {
		return nil, fmt.Errorf("The email address is required.")
	}
	if len(address) < 1 || 100 < len(address) {
		return nil, fmt.Errorf("Email address must be 100 characters or less.")
	}
	if !regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`).MatchString(address) {
		return nil, fmt.Errorf("Email address format is invalid.")
	}
	emailAddress.address = address
	return emailAddress, nil
}
