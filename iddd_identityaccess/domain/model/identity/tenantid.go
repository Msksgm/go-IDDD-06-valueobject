package identity

import (
	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
	"github.com/google/uuid"
)

type TenantId struct {
	id string
}

func NewTenantId(uu string) (*TenantId, error) {
	tenantId := new(TenantId)
	tenantId, err := setId(tenantId, uu)
	if err != nil {
		return nil, err
	}
	return tenantId, nil
}

func setId(tenantId *TenantId, uu string) (_ *TenantId, err error) {
	defer ierrors.Wrap(&err, "tenantid.setId(%s)", uu)
	if _, err := uuid.Parse(uu); err != nil {
		return nil, err
	}
	tenantId.id = uu
	return tenantId, nil
}
