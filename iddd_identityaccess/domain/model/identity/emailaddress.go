package identity

type EmailAddress struct {
	address string
}

func NewEmailAddress(address string) (_ *EmailAddress, err error) {
	emailAddress := new(EmailAddress)
	emailAddress.address = address
	return emailAddress, nil
}
