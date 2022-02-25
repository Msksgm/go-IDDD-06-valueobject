package contactinformation

import (
	"log"
	"testing"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/emailaddress"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/postaladdress"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/telephone"
	"github.com/google/go-cmp/cmp"
)

const (
	address                                                     = "sample@mail.com"
	streetAddress, city, stateProvince, postalCode, countryCode = "streetAddress", "city", "stateProvince", "postalCode", "00"
	primaryNumber                                               = "090-1234-5678"
	secondaryNumber                                             = "090-1234-5678"
)

var (
	emailAddress       *emailaddress.EmailAddress
	postalAddress      *postaladdress.PostalAddress
	primaryTelephone   *telephone.Telephone
	secondaryTelephone *telephone.Telephone
)

var err error

func init() {
	emailAddress, err = emailaddress.NewEmailAddress(address)
	if err != nil {
		log.Fatal(err)
	}
	postalAddress, err = postaladdress.NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
	if err != nil {
		log.Fatal(err)
	}
	primaryTelephone, err = telephone.NewTelephone(primaryNumber)
	if err != nil {
		log.Fatal(err)
	}
	secondaryTelephone, err = telephone.NewTelephone(secondaryNumber)
	if err != nil {
		log.Fatal(err)
	}
}

func TestNewContactInfomation(t *testing.T) {
	got, err := NewContactInformation(*emailAddress, *postalAddress, *primaryTelephone, *secondaryTelephone)
	if err != nil {
		t.Fatal(err)
	}
	want := &ContactInformation{emailAddress: *emailAddress, postalAddress: *postalAddress, primaryTelephone: *primaryTelephone, secondaryTelephone: *secondaryTelephone}

	allowUnexported := cmp.AllowUnexported(ContactInformation{}, emailaddress.EmailAddress{}, postaladdress.PostalAddress{}, telephone.Telephone{})
	if diff := cmp.Diff(want, got, allowUnexported); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
