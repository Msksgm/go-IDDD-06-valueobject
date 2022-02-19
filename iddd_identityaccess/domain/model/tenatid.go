package model

import (
	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
	"github.com/google/uuid"
)

type TenantId struct {
	tenantId string
}

func NewTenantId(uu string) (_ *TenantId, err error) {
	defer ierrors.Wrap(&err, "NewTenantId")
	if _, err := uuid.Parse(uu); err != nil {
		return nil, err
	}
	tenantId := &TenantId{tenantId: uu}
	return tenantId, nil
}
