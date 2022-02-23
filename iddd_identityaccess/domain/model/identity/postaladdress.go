package identity

import "fmt"

type PostalAddress struct {
	streetAddress string
	city          string
	stateProvince string
	postalCode    string
	countryCode   string
}

func NewPostalAddress(aStreetAddress string, aCity string, aStateProvince string, aPostalCode string, aCountryCode string) (_ *PostalAddress, err error) {
	postalAddress := new(PostalAddress)

	// set StreetAddress
	if aStreetAddress == "" {
		return nil, fmt.Errorf("The street address is required.")
	}
	if len(aStreetAddress) < 1 || 100 < len(aStreetAddress) {
		return nil, fmt.Errorf("The street address must be 100 characters or less.")
	}
	postalAddress.streetAddress = aStreetAddress
	postalAddress.city = aCity
	postalAddress.postalCode = aPostalCode
	postalAddress.stateProvince = aStateProvince
	postalAddress.countryCode = aCountryCode
	return postalAddress, nil
}
