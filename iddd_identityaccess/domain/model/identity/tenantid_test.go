package identity

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestNewTenantId(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenatId, err := NewTenantId(uu)
		if err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(tenatId.tenantId, uu) {
			t.Errorf("tenantId: %s should be equal to uu %s", tenatId, uu)
		}
	})
	t.Run("fail invalid UUID length", func(t *testing.T) {
		uu := "UUID"

		tenatId, err := NewTenantId(uu)
		want := fmt.Sprintf("tenantid.validateUu(%s): invalid UUID length: %d", uu, len(uu))
		if got := err.Error(); want != got {
			t.Errorf("got %s, want %s", got, want)
		}

		if tenatId != nil {
			t.Errorf("tenantId should be equal to nil, but %v", tenatId)
		}
	})
}
