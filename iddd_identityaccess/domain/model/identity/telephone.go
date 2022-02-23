package identity

import (
	"fmt"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
)

type Telephone struct {
	number string
}

func NewTelephone(aNumber string) (_ *Telephone, err error) {
	defer ierrors.Wrap(&err, "telephone.NewTelephone(%s)", aNumber)
	telephone := new(Telephone)
	if aNumber == "" {
		return nil, fmt.Errorf("Telephone number is required.")
	}
	if len(aNumber) < 5 || 20 < len(aNumber) {
		return nil, fmt.Errorf("Telephone number must be between 5 and 20 characters.")
	}
	telephone.number = aNumber
	return telephone, nil
}
