package identity

import (
	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
	"github.com/google/uuid"
)

type TenantId struct {
	tenantId string
}

func NewTenantId(uu string) (*TenantId, error) {
	if err := validateUu(uu); err != nil {
		return nil, err
	}
	tenantId := &TenantId{tenantId: uu}
	return tenantId, nil
}

func validateUu(uu string) (err error) {
	defer ierrors.Wrap(&err, "tenantid.validateUu(%s)", uu)
	if _, err := uuid.Parse(uu); err != nil {
		return err
	}
	return nil
}
