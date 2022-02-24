package contactinformation

import (
	"testing"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/emailaddress"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/postaladdress"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/telephone"
	"github.com/google/go-cmp/cmp"
)

func TestNewContactInfomation(t *testing.T) {
	address := "sample@mail.com"
	emailAddress, err := emailaddress.NewEmailAddress(address)
	if err != nil {
		t.Fatal(err)
	}
	streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", "city", "stateProvince", "postalCode", "00"
	postalAddress, err := postaladdress.NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
	if err != nil {
		t.Fatal(err)
	}
	primaryNumber := "090-1234-5678"
	primaryTelephone, err := telephone.NewTelephone(primaryNumber)
	if err != nil {
		t.Fatal(err)
	}
	secondaryNumber := "090-1234-5678"
	secondaryTlephone, err := telephone.NewTelephone(secondaryNumber)
	if err != nil {
		t.Fatal(err)
	}

	got, err := NewContactInformation(*emailAddress, *postalAddress, *primaryTelephone, *secondaryTlephone)
	if err != nil {
		t.Fatal(err)
	}
	want := &ContactInformation{emailAddress: *emailAddress, postalAddress: *postalAddress, primaryTelephone: *primaryTelephone, secondaryTelephone: *secondaryTlephone}

	allowUnexported := cmp.AllowUnexported(ContactInformation{}, emailaddress.EmailAddress{}, postaladdress.PostalAddress{}, telephone.Telephone{})
	if diff := cmp.Diff(want, got, allowUnexported); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
