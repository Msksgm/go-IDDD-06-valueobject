package model

import (
	"github.com/google/uuid"
)

type TenantId struct {
	tenantId string
}

func NewTenantId(uu string) (*TenantId, error) {
	if _, err := uuid.Parse(uu); err != nil {
		return nil, err
	}
	tenantId := &TenantId{tenantId: uu}
	return tenantId, nil
}
