package identity

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
)

type EmailAddress struct {
	address string
}

func NewEmailAddress(anAddress string) (_ *EmailAddress, err error) {
	defer ierrors.Wrap(&err, "emailaddress.NewEmailAddress(%s)", anAddress)
	emailAddress := new(EmailAddress)
	if anAddress == "" {
		return nil, fmt.Errorf("The email address is required.")
	}
	if len(anAddress) < 1 || 100 < len(anAddress) {
		return nil, fmt.Errorf("Email address must be 100 characters or less.")
	}
	if !regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`).MatchString(anAddress) {
		return nil, fmt.Errorf("Email address format is invalid.")
	}
	emailAddress.address = anAddress
	return emailAddress, nil
}

// TODO add the shallow copy of EmailAddress

func (emailAddress *EmailAddress) Address() string {
	return emailAddress.address
}

func (emailAddress *EmailAddress) Equals(otherEmailAddress EmailAddress) bool {
	return reflect.DeepEqual(emailAddress.address, otherEmailAddress.address)
}

func (emailAddress *EmailAddress) String() string {
	return fmt.Sprintf("EmailAddress [address=%s]", emailAddress.address)
}
