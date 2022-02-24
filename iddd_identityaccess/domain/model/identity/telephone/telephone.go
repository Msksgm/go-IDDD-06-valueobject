package telephone

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
)

type Telephone struct {
	number string
}

func NewTelephone(aNumber string) (_ *Telephone, err error) {
	defer ierrors.Wrap(&err, "telephone.NewTelephone(%s)", aNumber)
	telephone := new(Telephone)

	// set number
	if aNumber == "" {
		return nil, fmt.Errorf("Telephone number is required.")
	}
	if len(aNumber) < 5 || 20 < len(aNumber) {
		return nil, fmt.Errorf("Telephone number must be between 5 and 20 characters.")
	}
	if !regexp.MustCompile(`^0\d{2,3}-\d{1,4}-\d{4}$`).MatchString(aNumber) {
		return nil, fmt.Errorf("Telephone number or its format is invalid.")
	}
	telephone.number = aNumber

	return telephone, nil
}

// TODO add the shallow copy function

func (telephone *Telephone) Number() string {
	return telephone.number
}

func (telephone *Telephone) Equal(otherTelephone *Telephone) bool {
	return reflect.DeepEqual(telephone.number, otherTelephone.number)
}

func (telephone *Telephone) String() string {
	return fmt.Sprintf("Telephone [number=%v]", telephone.number)
}
