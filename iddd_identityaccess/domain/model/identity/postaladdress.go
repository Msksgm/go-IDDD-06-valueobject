package identity

import (
	"fmt"
	"reflect"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
)

type PostalAddress struct {
	streetAddress string
	city          string
	stateProvince string
	postalCode    string
	countryCode   string
}

func NewPostalAddress(aStreetAddress string, aCity string, aStateProvince string, aPostalCode string, aCountryCode string) (_ *PostalAddress, err error) {
	defer ierrors.Wrap(&err, "postaladdress.NewPostalAddress(%s, %s, %s, %s, %s)", aStreetAddress, aCity, aStateProvince, aPostalCode, aCountryCode)

	// validate StreetAddress
	if err := ierrors.NewArgumentNotEmptyError(aStreetAddress, "The street address is required.").GetError(); err != nil {
		return nil, err
	}
	if err := ierrors.NewArgumentLengthError(aStreetAddress, 1, 100, "The street address must be 100 characters or less.").GetError(); err != nil {
		return nil, err
	}

	// validate City
	if err := ierrors.NewArgumentNotEmptyError(aCity, "The street city is required.").GetError(); err != nil {
		return nil, err
	}
	if err := ierrors.NewArgumentLengthError(aCity, 1, 100, "The street city must be 100 characters or less.").GetError(); err != nil {
		return nil, err
	}

	// validate PostalCode
	if err := ierrors.NewArgumentNotEmptyError(aPostalCode, "The postal code is required.").GetError(); err != nil {
		return nil, err
	}
	if err := ierrors.NewArgumentLengthError(aPostalCode, 5, 12, "The postal code must be between 5 characters and 12 characters.").GetError(); err != nil {
		return nil, err
	}

	// validate StateProvince
	if err := ierrors.NewArgumentNotEmptyError(aStateProvince, "The state/province is required.").GetError(); err != nil {
		return nil, err
	}
	if err := ierrors.NewArgumentLengthError(aStateProvince, 2, 100, "The state/province must be between 2 charcters and 100 characters").GetError(); err != nil {
		return nil, err
	}

	// validate CountryCode
	if err := ierrors.NewArgumentNotEmptyError(aCountryCode, "The country is required.").GetError(); err != nil {
		return nil, err
	}
	if err := ierrors.NewArgumentLengthError(aCountryCode, 2, 2, "The country code must be two characters.").GetError(); err != nil {
		return nil, err
	}

	return &PostalAddress{streetAddress: aStreetAddress, city: aCity, postalCode: aPostalCode, stateProvince: aStateProvince, countryCode: aCountryCode}, nil
}

// TODO add shallow copy of PostalAddress

func (postalAddress *PostalAddress) City() string {
	return postalAddress.city
}

func (postalAddress *PostalAddress) CountryCode() string {
	return postalAddress.countryCode
}

func (postalAddress *PostalAddress) PostalCode() string {
	return postalAddress.postalCode
}

func (postalAddress *PostalAddress) StateProvince() string {
	return postalAddress.stateProvince
}

func (postalAddress *PostalAddress) StreetAddress() string {
	return postalAddress.streetAddress
}

func (postalAddress *PostalAddress) Equals(otherPostalAddress PostalAddress) bool {
	isStreetAddressEqual := reflect.DeepEqual(postalAddress.StreetAddress(), otherPostalAddress.StreetAddress())
	isCityEqual := reflect.DeepEqual(postalAddress.City(), otherPostalAddress.City())
	isStateProvinceEqual := reflect.DeepEqual(postalAddress.StateProvince(), otherPostalAddress.StateProvince())
	isCountryCodeEqual := reflect.DeepEqual(postalAddress.CountryCode(), otherPostalAddress.CountryCode())
	isPostalCodeEqual := reflect.DeepEqual(postalAddress.PostalCode(), otherPostalAddress.PostalCode())
	return isStreetAddressEqual && isCityEqual && isStateProvinceEqual && isCountryCodeEqual && isPostalCodeEqual
}

func (postaladdress *PostalAddress) String() string {
	return fmt.Sprintf("PostalAddress [streetAddress=%s, city=%s, stateProvince=%s, postalCode=%s, countryCode=%s]", postaladdress.streetAddress, postaladdress.city, postaladdress.stateProvince, postaladdress.postalCode, postaladdress.countryCode)
}
