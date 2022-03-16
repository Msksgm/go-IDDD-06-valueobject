package identity

type Person struct {
	tenantId           TenantId
	name               FullName
	contactInformation ContactInformation
}

func NewPerson(aTenantId TenantId, aName FullName, aContactInformation ContactInformation) (*Person, error) {
	person := new(Person)

	person.tenantId = aTenantId
	person.name = aName
	person.contactInformation = aContactInformation
	return person, nil
}

func (person *Person) ChangeContactInformation(aContactInformation ContactInformation) error {
	person.contactInformation = aContactInformation
	return nil
}

func (person *Person) ChangeName(aName FullName) error {
	person.name = aName
	return nil
}

func (person *Person) ContactInformation() ContactInformation {
	return person.contactInformation
}

func (person *Person) EmailAddress() EmailAddress {
	return *person.contactInformation.EmailAddress()
}

func (person *Person) Name() FullName {
	return person.name
}
