package identity

type PostalAddress struct {
	streetAddress string
	city          string
	stateProvince string
	postalCode    string
	countryCode   string
}

func NewPostalAddress(aStreetAddress string, aCity string, aStateProvince string, aPostalCode string, aCountryCode string) (_ *PostalAddress, err error) {
	postalAddress := new(PostalAddress)
	postalAddress.streetAddress = aStreetAddress
	postalAddress.city = aCity
	postalAddress.stateProvince = aStateProvince
	postalAddress.countryCode = aCountryCode
	return postalAddress, nil
}
