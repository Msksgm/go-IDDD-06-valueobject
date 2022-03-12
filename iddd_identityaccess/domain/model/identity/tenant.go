package identity

import (
	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
)

type Tenant struct {
	tenantId TenantId
	name     string
	active   bool
}

func NewTenant(aTenantId TenantId, aName string, anActive bool) (_ *Tenant, err error) {
	defer ierrors.Wrap(&err, "tenant.NewTenant(%v, %v, %v)", aTenantId, aName, anActive)
	// validate name
	if err := ierrors.NewArgumentNotEmptyError(aName, "The tenant name is required.").GetError(); err != nil {
		return nil, err
	}
	if err := ierrors.NewArgumentLengthError(aName, 1, 100, "The tenant description must be 100 characters or less.").GetError(); err != nil {
		return nil, err
	}

	return &Tenant{tenantId: aTenantId, name: aName, active: anActive}, nil
}

func (tenant *Tenant) setActive(active bool) {
	tenant.active = active
}

func (tenant *Tenant) Activate() {
	if !tenant.IsActive() {
		tenant.setActive(true)
	}
}

func (tenant *Tenant) Deactivate() {
	if tenant.IsActive() {
		tenant.setActive(false)
	}
}

func (tenant *Tenant) IsActive() bool {
	return tenant.active
}

func (tenant *Tenant) Equals(otherTenant Tenant) bool {
	return tenant.tenantId == otherTenant.tenantId
}
