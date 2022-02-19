package identity

import "fmt"

type Tenant struct {
	tenantId TenantId
	name     string
}

func NewTenant(tenantId TenantId, name string) (_ *Tenant, err error) {
	if name == "" {
		return nil, fmt.Errorf("The tenant name is required.")
	}
	if len(name) < 1 || len(name) > 100 {
		return nil, fmt.Errorf("The tenant description must be 100 characters or less.")
	}
	return &Tenant{tenantId: tenantId, name: name}, nil
}
