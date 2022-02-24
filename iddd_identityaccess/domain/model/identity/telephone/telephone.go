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
	if err := ierrors.NewArgumentNotEmptyError(aNumber, "Telephone number is required.").GetError(); err != nil {
		return nil, err
	}
	if err := ierrors.NewArgumentLengthError(aNumber, 5, 20, "Telephone number must be between 5 and 20 characters.").GetError(); err != nil {
		return nil, err
	}
	if err := ierrors.NewArgumentTrueErrorArguments(regexp.MustCompile(`^0\d{2,3}-\d{1,4}-\d{4}$`).MatchString(aNumber), "First name must be at least one character in length, starting with a capital letter.").GetError(); err != nil {
		return nil, err
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
