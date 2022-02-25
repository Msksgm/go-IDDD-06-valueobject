package contactinformation

import (
	"log"
	"testing"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/emailaddress"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/postaladdress"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/telephone"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

const (
	address                                                     = "sample@mail.com"
	streetAddress, city, stateProvince, postalCode, countryCode = "streetAddress", "city", "stateProvince", "postalCode", "00"
	primaryNumber                                               = "090-1234-5678"
	secondaryNumber                                             = "090-5678-1234"
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

func TestChangeEmailAddress(t *testing.T) {
	contactInformation, err := NewContactInformation(*emailAddress, *postalAddress, *primaryTelephone, *secondaryTelephone)
	if err != nil {
		t.Fatal(err)
	}
	copiedContactInformation, err := CopyContactInfomation(*contactInformation)
	if err != nil {
		t.Fatal(err)
	}

	changedEmailAddress, err := emailaddress.NewEmailAddress("changed@email.com")
	if err != nil {
		t.Fatal(err)
	}
	contactInformation2, err := contactInformation.ChangeEmailAddress(*changedEmailAddress)
	if err != nil {
		t.Fatal(err)
	}

	if !(*contactInformation == *copiedContactInformation) {
		t.Fatalf("contactInfomation: %v must be equal to copiedContactInformation: %v", contactInformation, copiedContactInformation)
	}

	if contactInformation.Equals(*contactInformation2) {
		t.Fatalf("contactInfomation: %v must not be equal to copiedContactInformation: %v", contactInformation, contactInformation2)
	}

	if copiedContactInformation.Equals(*contactInformation2) {
		t.Fatalf("contactInfomationj: %v must not be equal to copiedContactInformation: %v", copiedContactInformation, contactInformation2)
	}

	if contactInformation2.EmailAddress().Address() != "changed@email.com" {
		t.Fatalf("contactInformation2.EmailAddress().Address(): %v must be equal to copiedContactInformation: %v", contactInformation2.EmailAddress().Address(), changedEmailAddress)
	}

	opts := cmp.Options{
		cmp.AllowUnexported(ContactInformation{}, emailaddress.EmailAddress{}, postaladdress.PostalAddress{}, telephone.Telephone{}),
		cmpopts.IgnoreFields(emailaddress.EmailAddress{}),
	}
	if diff := cmp.Diff(contactInformation, copiedContactInformation, opts); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestPostalAddress(t *testing.T) {
	contactInformation, err := NewContactInformation(*emailAddress, *postalAddress, *primaryTelephone, *secondaryTelephone)
	if err != nil {
		t.Fatal(err)
	}
	copiedContactInformation, err := CopyContactInfomation(*contactInformation)
	if err != nil {
		t.Fatal(err)
	}

	changedPostalAddress, err := postaladdress.NewPostalAddress("changedStreetAddress", "changedCity", "changedStateProvince", "postalCode", "00")
	if err != nil {
		t.Fatal(err)
	}
	contactInformation2, err := contactInformation.ChangePostalAddress(*changedPostalAddress)
	if err != nil {
		t.Fatal(err)
	}

	if !(*contactInformation == *copiedContactInformation) {
		t.Fatalf("contactInfomation: %v must be equal to copiedContactInformation: %v", contactInformation, copiedContactInformation)
	}

	if contactInformation.Equals(*contactInformation2) {
		t.Fatalf("contactInfomation: %v must not be equal to copiedContactInformation: %v", contactInformation, contactInformation2)
	}

	if copiedContactInformation.Equals(*contactInformation2) {
		t.Fatalf("contactInfomation: %v must not be equal to copiedContactInformation: %v", copiedContactInformation, contactInformation2)
	}

	if contactInformation2.PostalAddress().City() != changedPostalAddress.City() {
		t.Fatalf("contactInformation2.PostalAddress().City(): %v must not be equal to copiedContactInformation: %v", contactInformation2.PostalAddress().City(), changedPostalAddress.City())
	}

	if contactInformation2.PostalAddress().CountryCode() != changedPostalAddress.CountryCode() {
		t.Fatalf("contactInformation2.PostalAddress().CountryCode(): %v must not be equal to copiedContactInformation: %v", contactInformation2.PostalAddress().CountryCode(), changedPostalAddress.CountryCode())
	}

	if contactInformation2.PostalAddress().PostalCode() != changedPostalAddress.PostalCode() {
		t.Fatalf("contactInformation2.PostalAddress().PostalCode(): %v must not be equal to copiedContactInformation: %v", contactInformation2.PostalAddress().PostalCode(), changedPostalAddress.PostalCode())
	}

	if contactInformation2.PostalAddress().StateProvince() != changedPostalAddress.StateProvince() {
		t.Fatalf("contactInformation2.PostalAddress().StateProvince(): %v must not be equal to copiedContactInformation: %v", contactInformation2.PostalAddress().StateProvince(), changedPostalAddress.StateProvince())
	}

	if contactInformation2.PostalAddress().StreetAddress() != changedPostalAddress.StreetAddress() {
		t.Fatalf("contactInformation2.PostalAddress().StreetAddress(): %v must be equal to copiedContactInformation: %v", contactInformation2.PostalAddress().StreetAddress(), changedPostalAddress.StreetAddress())
	}

	opts := cmp.Options{
		cmp.AllowUnexported(ContactInformation{}, emailaddress.EmailAddress{}, postaladdress.PostalAddress{}, telephone.Telephone{}),
		cmpopts.IgnoreFields(postaladdress.PostalAddress{}),
	}
	if diff := cmp.Diff(contactInformation, copiedContactInformation, opts); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestPrimaryTelephone(t *testing.T) {
	contactInformation, err := NewContactInformation(*emailAddress, *postalAddress, *primaryTelephone, *secondaryTelephone)
	if err != nil {
		t.Fatal(err)
	}
	copiedContactInformation, err := CopyContactInfomation(*contactInformation)
	if err != nil {
		t.Fatal(err)
	}

	changedPrimaryTelephone, err := telephone.NewTelephone("090-0000-0000")
	if err != nil {
		t.Fatal(err)
	}
	contactInformation2, err := contactInformation.ChangePrimaryTelephone(*changedPrimaryTelephone)
	if err != nil {
		t.Fatal(err)
	}

	if !(*contactInformation == *copiedContactInformation) {
		t.Fatalf("contactInfomation: %v must be equal to copiedContactInformation: %v", contactInformation, copiedContactInformation)
	}

	if contactInformation.Equals(*contactInformation2) {
		t.Fatalf("contactInfomation: %v must not be equal to copiedContactInformation: %v", contactInformation, contactInformation2)
	}

	if copiedContactInformation.Equals(*contactInformation2) {
		t.Fatalf("contactInfomation: %v must not be equal to copiedContactInformation: %v", copiedContactInformation, contactInformation2)
	}

	if contactInformation2.PrimaryTelephone().Number() != changedPrimaryTelephone.Number() {
		t.Fatalf("contactInformation2.PrimaryTelephone().Number(): %v must be equal to copiedContactInformation: %v", contactInformation2.PrimaryTelephone().Number(), changedPrimaryTelephone.Number())
	}

	opts := cmp.Options{
		cmp.AllowUnexported(ContactInformation{}, emailaddress.EmailAddress{}, postaladdress.PostalAddress{}, telephone.Telephone{}),
		cmpopts.IgnoreFields(telephone.Telephone{}),
	}
	if diff := cmp.Diff(contactInformation, copiedContactInformation, opts); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
