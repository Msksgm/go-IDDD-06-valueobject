package identity

import (
	"fmt"

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
	emailAddress.address = address
	return emailAddress, nil
}
