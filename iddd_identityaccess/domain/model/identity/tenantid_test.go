package identity

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewTenantId(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		got, err := NewTenantId(uuidV4)
		if err != nil {
			t.Fatal(err)
		}

		want := &TenantId{id: uuidV4}

		if diff := cmp.Diff(want, got, cmp.AllowUnexported(TenantId{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail invalid UUID length", func(t *testing.T) {
		dumpyUuuid := "UUID"

		tenatId, err := NewTenantId(dumpyUuuid)
		want := fmt.Sprintf("tenantid.NewTenantId(%s): invalid UUID length: %d", dumpyUuuid, len(dumpyUuuid))
		if got := err.Error(); want != got {
			t.Errorf("got %s, want %s", got, want)
		}

		if tenatId != nil {
			t.Errorf("tenantId should be equal to nil, but %v", tenatId)
		}
	})
}

func TestTenantIdEquals(t *testing.T) {
	tenantId, err := NewTenantId(uuidV4)
	if err != nil {
		t.Fatal(err)
	}

	otherTenantId := &TenantId{id: uuidV4}

	if !tenantId.Equals(otherTenantId) {
		t.Errorf("tenantId: %v must be euqal to otherTenantId %v", tenantId, otherTenantId)
	}
}
