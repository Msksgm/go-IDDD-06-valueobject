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

	// set PostalCode
	if aPostalCode == "" {
		return nil, fmt.Errorf("The postal code is required.")
	}
	if len(aPostalCode) < 5 || 12 < len(aPostalCode) {
		return nil, fmt.Errorf("The postal code must be between 5 characters and 12 characters.")
	}
	postalAddress.postalCode = aPostalCode

	// set StateProvince
	if aStateProvince == "" {
		return nil, fmt.Errorf("The state/province is required.")
	}
	if len(aStateProvince) < 2 || 100 < len(aStateProvince) {
		return nil, fmt.Errorf("The state/province must be between 2 charcters and 100 characters")
	}
	postalAddress.stateProvince = aStateProvince

	// set CountryCode
	if aCountryCode == "" {
		return nil, fmt.Errorf("The country is required.")
	}
	if len(aCountryCode) < 2 || 2 < len(aCountryCode) {
		return nil, fmt.Errorf("The country code must be two characters.")
	}
	postalAddress.countryCode = aCountryCode
	return postalAddress, nil
}

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
