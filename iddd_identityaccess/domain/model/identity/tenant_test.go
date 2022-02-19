package identity

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var rs1Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	return string(b)
}

func TestNewTenant(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenantId, err := NewTenantId(uu)
		if err != nil {
			t.Fatal(err)
		}

		name := "TenantName"
		got, err := NewTenant(*tenantId, name)
		if err != nil {
			t.Fatal(err)
		}

		want := &Tenant{tenantId: TenantId{id: uu}, name: "TenantName"}

		if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tenant{}, TenantId{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail empty name", func(t *testing.T) {
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenantId, err := NewTenantId(uu)
		if err != nil {
			t.Fatal(err)
		}

		name := ""
		tenant, err := NewTenant(*tenantId, name)
		want := fmt.Sprintf("tenant.setName(%s): The tenant name is required.", name)
		if got := err.Error(); want != got {
			t.Errorf("got %s, want %s", got, want)
		}
		if tenant != nil {
			t.Errorf("tenant should be nil, but %v", tenant)
		}
	})
	t.Run("fail over 100 characters name", func(t *testing.T) {
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenantId, err := NewTenantId(uu)
		if err != nil {
			t.Fatal(err)
		}

		name := RandString(101)
		tenant, err := NewTenant(*tenantId, name)
		want := fmt.Sprintf("tenant.setName(%s): The tenant description must be 100 characters or less.", name)
		if got := err.Error(); want != got {
			t.Errorf("got %s, want %s", got, want)
		}
		if tenant != nil {
			t.Errorf("tenant should be nil, but %v", tenant)
		}
	})
}

func TestEquals(t *testing.T) {
	u, err := uuid.NewRandom()
	if err != nil {
		t.Fatal(err)
	}
	uu := u.String()

	tenantId, err := NewTenantId(uu)
	if err != nil {
		t.Fatal(err)
	}

	name1 := "TenantName1"
	tenant1, err := NewTenant(*tenantId, name1)
	if err != nil {
		t.Fatal(err)
	}

	name2 := "TenantName1"
	tenant2, err := NewTenant(*tenantId, name2)
	if err != nil {
		t.Fatal(err)
	}

	if !tenant1.Equals(*tenant2) {
		t.Errorf("tenant1 %v must be equal to %v by tenantId", tenant1, tenant2)
	}
}
