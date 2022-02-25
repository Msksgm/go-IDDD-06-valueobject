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

func CopyContactInfomation(aContactInformation ContactInformation) (*ContactInformation, error) {
	copiedContactInformation, err := NewContactInformation(aContactInformation.emailAddress, aContactInformation.postalAddress, aContactInformation.primaryTelephone, aContactInformation.secondaryTelephone)
	if err != nil {
		return nil, err
	}
	return copiedContactInformation, nil
}

func (contactInformation *ContactInformation) ChangeEmailAddress(aEmailaddress emailaddress.EmailAddress) (*ContactInformation, error) {
	newContactInformation, err := NewContactInformation(aEmailaddress, contactInformation.postalAddress, contactInformation.primaryTelephone, contactInformation.secondaryTelephone)
	if err != nil {
		return nil, err
	}
	return newContactInformation, nil
}

func (contactInformation *ContactInformation) ChangePostalAddress(aPostalAddress postaladdress.PostalAddress) (*ContactInformation, error) {
	newContactInformation, err := NewContactInformation(contactInformation.emailAddress, aPostalAddress, contactInformation.primaryTelephone, contactInformation.secondaryTelephone)
	if err != nil {
		return nil, err
	}
	return newContactInformation, nil
}

func (contactInformation *ContactInformation) ChangePrimaryTelephone(aTelephone telephone.Telephone) (*ContactInformation, error) {
	newContactInformation, err := NewContactInformation(contactInformation.emailAddress, contactInformation.postalAddress, aTelephone, contactInformation.secondaryTelephone)
	if err != nil {
		return nil, err
	}
	return newContactInformation, nil
}

func (contactInformation *ContactInformation) ChangeSecondaryTelephone(aTelephone telephone.Telephone) (*ContactInformation, error) {
	newContactInformation, err := NewContactInformation(contactInformation.emailAddress, contactInformation.postalAddress, contactInformation.primaryTelephone, aTelephone)
	if err != nil {
		return nil, err
	}
	return newContactInformation, nil
}

func (contactInformation *ContactInformation) EmailAddress() *emailaddress.EmailAddress {
	return &contactInformation.emailAddress
}

func (contactInformation *ContactInformation) PostalAddress() *postaladdress.PostalAddress {
	return &contactInformation.postalAddress
}

func (contactInformation *ContactInformation) PrimaryTelephone() *telephone.Telephone {
	return &contactInformation.primaryTelephone
}

func (contactInformation *ContactInformation) SecondaryTelephone() *telephone.Telephone {
	return &contactInformation.secondaryTelephone
}

func (contactInformation ContactInformation) Equals(otheContactInformation ContactInformation) bool {
	isEmailAddressEqual := contactInformation.emailAddress.Equals(otheContactInformation.emailAddress)
	isPostalAddressEqual := contactInformation.postalAddress.Equals(otheContactInformation.postalAddress)
	isPrimaryTelephoneEqual := contactInformation.primaryTelephone.Equals(otheContactInformation.primaryTelephone)
	isSecondaryTelephoneEqual := contactInformation.secondaryTelephone.Equals(otheContactInformation.secondaryTelephone)
	return isEmailAddressEqual && isPostalAddressEqual && isPrimaryTelephoneEqual && isSecondaryTelephoneEqual
}
