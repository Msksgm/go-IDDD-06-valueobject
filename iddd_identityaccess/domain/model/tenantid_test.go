package model

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestNewTenantId(t *testing.T) {
	u, err := uuid.NewRandom()
	if err != nil {
		t.Error(err)
	}
	uu := u.String()

	tenatId, err := NewTenantId(uu)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(tenatId.tenantId, uu) {
		t.Errorf("tenantId: %s is should be equal to uu %s", tenatId, uu)
	}
}
