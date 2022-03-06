package identity

import (
	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
	"github.com/google/uuid"
)

type TenantId struct {
	id string
}

func NewTenantId(uu string) (_ *TenantId, err error) {
	defer ierrors.Wrap(&err, "tenantid.NewTenantId(%s)", uu)

	// setId
	if _, err := uuid.Parse(uu); err != nil {
		return nil, err
	}

	return &TenantId{id: uu}, nil
}

func (tenantId *TenantId) Equals(otherTeanntId *TenantId) bool {
	return tenantId.id == otherTeanntId.id
}
