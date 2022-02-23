package identity

import (
	"fmt"

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
	postalAddress := new(PostalAddress)

	// set StreetAddress
	if aStreetAddress == "" {
		return nil, fmt.Errorf("The street address is required.")
	}
	if len(aStreetAddress) < 1 || 100 < len(aStreetAddress) {
		return nil, fmt.Errorf("The street address must be 100 characters or less.")
	}
	postalAddress.streetAddress = aStreetAddress

	// set City
	if aCity == "" {
		return nil, fmt.Errorf("The street city is required.")
	}
	if len(aCity) < 1 || 100 < len(aCity) {
		return nil, fmt.Errorf("The street city must be 100 characters or less.")
	}
	postalAddress.city = aCity

	postalAddress.postalCode = aPostalCode
	postalAddress.stateProvince = aStateProvince
	postalAddress.countryCode = aCountryCode
	return postalAddress, nil
}
