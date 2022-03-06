package identity

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func TestNewTenantId(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		got, err := NewTenantId(uu)
		if err != nil {
			t.Fatal(err)
		}

		want := &TenantId{id: uu}

		if diff := cmp.Diff(want, got, cmp.AllowUnexported(TenantId{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail invalid UUID length", func(t *testing.T) {
		uu := "UUID"

		tenatId, err := NewTenantId(uu)
		want := fmt.Sprintf("tenantid.NewTenantId(%s): invalid UUID length: %d", uu, len(uu))
		if got := err.Error(); want != got {
			t.Errorf("got %s, want %s", got, want)
		}

		if tenatId != nil {
			t.Errorf("tenantId should be equal to nil, but %v", tenatId)
		}
	})
}

func TestTenantIdEquals(t *testing.T) {
	u, err := uuid.NewRandom()
	if err != nil {
		t.Fatal(err)
	}
	uu := u.String()

	tenantId, err := NewTenantId(uu)
	if err != nil {
		t.Fatal(err)
	}

	otherTenantId := &TenantId{id: uu}

	if !tenantId.Equals(otherTenantId) {
		t.Errorf("tenantId: %v must be euqal to otherTenantId %v", tenantId, otherTenantId)
	}
}
