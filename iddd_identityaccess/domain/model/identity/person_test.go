package identity

import (
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

const (
	firstName = "FirstName"
	lastName  = "lastName"
)

var (
	contactInformation *ContactInformation
	fullName           *FullName
)

func init() {
	emailAddress, err := NewEmailAddress(address)
	if err != nil {
		log.Fatal(err)
	}
	postalAddress, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
	if err != nil {
		log.Fatal(err)
	}
	primaryTelephone, err := NewTelephone(primaryNumber)
	if err != nil {
		log.Fatal(err)
	}
	secondaryTelephone, err := NewTelephone(secondaryNumber)
	if err != nil {
		log.Fatal(err)
	}
	contactInformation, err = NewContactInformation(*emailAddress, *postalAddress, *primaryTelephone, *secondaryTelephone)
	if err != nil {
		log.Fatal(err)
	}

	fullName, err = NewFullName(firstName, lastName)
	if err != nil {
		log.Fatal(err)
	}

	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}
	uuString := uuid.String()
	tenantId, err = NewTenantId(uuString)
	if err != nil {
		log.Fatal(err)
	}
}

func TestNewPerson(t *testing.T) {
	got, err := NewPerson(*tenantId, *fullName, *contactInformation)
	if err != nil {
		t.Fatal(err)
	}

	want := &Person{tenantId: *tenantId, name: *fullName, contactInformation: *contactInformation}

	allowUnexported := cmp.AllowUnexported(Person{}, TenantId{}, FullName{}, ContactInformation{}, EmailAddress{}, PostalAddress{}, Telephone{})
	if diff := cmp.Diff(want, got, allowUnexported); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestChangeContactInformation(t *testing.T) {
	person, err := NewPerson(*tenantId, *fullName, *contactInformation)
	if err != nil {
		t.Fatal(err)
	}

	changedAddressString := "changed@email.com"
	changedAddress, err := NewEmailAddress(changedAddressString)
	if err != nil {
		t.Fatal(err)
	}
	changedContactInfomation, err := contactInformation.ChangeEmailAddress(*changedAddress)
	if err != nil {
		t.Fatal(err)
	}
	want, err := NewPerson(*tenantId, *fullName, *changedContactInfomation)
	if err != nil {
		t.Fatal(err)
	}
	if err := person.ChangeContactInformation(*changedContactInfomation); err != nil {
		t.Fatal(err)
	}

	allowUnexported := cmp.AllowUnexported(Person{}, TenantId{}, FullName{}, ContactInformation{}, EmailAddress{}, PostalAddress{}, Telephone{})
	if diff := cmp.Diff(want, person, allowUnexported); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestChangeName(t *testing.T) {
	person, err := NewPerson(*tenantId, *fullName, *contactInformation)
	if err != nil {
		t.Fatal(err)
	}

	changedFirstName := "ChangedFirstName"
	changedFullName, err := fullName.WithChangedFirstName(changedFirstName)
	if err != nil {
		t.Fatal(err)
	}

	if err := person.ChangeName(*changedFullName); err != nil {
		t.Fatal(err)
	}

	want, err := NewPerson(*tenantId, *changedFullName, *contactInformation)
	if err != nil {
		t.Fatal(err)
	}

	allowUnexported := cmp.AllowUnexported(Person{}, TenantId{}, FullName{}, ContactInformation{}, EmailAddress{}, PostalAddress{}, Telephone{})
	if diff := cmp.Diff(want, person, allowUnexported); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestContactInfomation(t *testing.T) {
	person, err := NewPerson(*tenantId, *fullName, *contactInformation)
	if err != nil {
		t.Fatal(err)
	}
	want := *contactInformation
	got := person.ContactInformation()
	allowUnexported := cmp.AllowUnexported(Person{}, TenantId{}, FullName{}, ContactInformation{}, EmailAddress{}, PostalAddress{}, Telephone{})
	if diff := cmp.Diff(want, got, allowUnexported); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestEmailAddress(t *testing.T) {
	person, err := NewPerson(*tenantId, *fullName, *contactInformation)
	if err != nil {
		t.Fatal(err)
	}
	want := *contactInformation.EmailAddress()
	got := person.EmailAddress()
	allowUnexported := cmp.AllowUnexported(Person{}, TenantId{}, FullName{}, ContactInformation{}, EmailAddress{}, PostalAddress{}, Telephone{})
	if diff := cmp.Diff(want, got, allowUnexported); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestName(t *testing.T) {
	person, err := NewPerson(*tenantId, *fullName, *contactInformation)
	if err != nil {
		t.Fatal(err)
	}
	want := *fullName
	got := person.Name()
	allowUnexported := cmp.AllowUnexported(Person{}, TenantId{}, FullName{}, ContactInformation{}, EmailAddress{}, PostalAddress{}, Telephone{})
	if diff := cmp.Diff(want, got, allowUnexported); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
