package contactinformation

import (
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/emailaddress"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/postaladdress"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/telephone"
)

type ContactInformation struct {
	emailAddress       emailaddress.EmailAddress
	postalAddress      postaladdress.PostalAddress
	primaryTelephone   telephone.Telephone
	secondaryTelephone telephone.Telephone
}

func NewContactInformation(aEmailaddress emailaddress.EmailAddress, aPostalAddress postaladdress.PostalAddress, aPrimaryAddress telephone.Telephone, aSecondaryAddress telephone.Telephone) (*ContactInformation, error) {
	contactInformation := new(ContactInformation)

	contactInformation.emailAddress = aEmailaddress
	contactInformation.postalAddress = aPostalAddress
	contactInformation.primaryTelephone = aPrimaryAddress
	contactInformation.secondaryTelephone = aSecondaryAddress

	return contactInformation, nil
}
