package emailaddress

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
	if err := ierrors.NewArgumentNotEmptyError(anAddress, "The email address is required.").GetError(); err != nil {
		return nil, err
	}
	if err := ierrors.NewArgumentLengthError(anAddress, 1, 100, "Email address must be 100 characters or less.").GetError(); err != nil {
		return nil, err
	}
	if err := ierrors.NewArgumentTrueErrorArguments(regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`).MatchString(anAddress), "Email address format is invalid.").GetError(); err != nil {
		return nil, err
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
