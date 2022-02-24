package tenantid

import (
	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
	"github.com/google/uuid"
)

type TenantId struct {
	id string
}

func NewTenantId(uu string) (_ *TenantId, err error) {
	defer ierrors.Wrap(&err, "tenantid.NewTenantId(%s)", uu)
	tenantId := new(TenantId)

	// setId
	if _, err := uuid.Parse(uu); err != nil {
		return nil, err
	}
	tenantId.id = uu

	return tenantId, nil
}
