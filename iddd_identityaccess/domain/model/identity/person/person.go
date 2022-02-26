package person

import (
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/contactinformation"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/fullname"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/tenantid"
)

type Person struct {
	tenantId           tenantid.TenantId
	name               fullname.FullName
	contactInformation contactinformation.ContactInformation
}

func NewPerson(aTenantId tenantid.TenantId, aName fullname.FullName, aContactInformation contactinformation.ContactInformation) (*Person, error) {
	person := new(Person)

	person.tenantId = aTenantId
	person.name = aName
	person.contactInformation = aContactInformation
	return person, nil
}

func (person *Person) ChangeContactInformation(aContactInformation contactinformation.ContactInformation) error {
	person.contactInformation = aContactInformation
	return nil
}
