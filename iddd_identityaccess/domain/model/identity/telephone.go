package identity

type Telephone struct {
	number string
}

func NewTelephone(aNumber string) (_ *Telephone, err error) {
	telephone := new(Telephone)
	telephone.number = aNumber
	return telephone, nil
}
