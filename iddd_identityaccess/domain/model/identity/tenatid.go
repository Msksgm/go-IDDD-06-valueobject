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
	if err := tenantId.setId(uu); err != nil {
		return nil, err
	}
	return tenantId, nil
}

func (tenantId *TenantId) setId(uu string) (err error) {
	defer ierrors.Wrap(&err, "tenantid.setId(%s)", uu)
	if _, err := uuid.Parse(uu); err != nil {
		return err
	}
	tenantId.id = uu
	return nil
}
