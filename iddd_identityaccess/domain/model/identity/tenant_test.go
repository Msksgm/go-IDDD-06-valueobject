package identity

import (
	"fmt"
	"testing"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/utils"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

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
		got, err := NewTenant(*tenantId, name, true)
		if err != nil {
			t.Fatal(err)
		}

		want := &Tenant{tenantId: TenantId{id: uu}, name: "TenantName", active: true}

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
		active := true
		tenant, err := NewTenant(*tenantId, name, active)
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

		name := utils.RandString(101)
		active := true
		tenant, err := NewTenant(*tenantId, name, active)
		want := fmt.Sprintf("tenant.setName(%s): The tenant description must be 100 characters or less.", name)
		if got := err.Error(); want != got {
			t.Errorf("got %s, want %s", got, want)
		}
		if tenant != nil {
			t.Errorf("tenant should be nil, but %v", tenant)
		}
	})
	t.Run("fail empty tenantId", func(t *testing.T) {
		tenantId := TenantId{id: ""}
		name := "TenantName"
		active := true
		tenant, err := NewTenant(tenantId, name, active)
		want := fmt.Sprintf("tenant.setTenantId(%s): TenentId is required.", tenantId)
		if got := err.Error(); want != got {
			t.Errorf("got %s, want %s", got, want)
		}
		if tenant != nil {
			t.Errorf("tenant should be nil, but %v", tenant)
		}
	})
}

func TestDeactivate(t *testing.T) {
	t.Run("active to deactive", func(t *testing.T) {
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenantId := TenantId{id: uu}
		name := "TenantName"
		acitve := true
		tenant := Tenant{tenantId: tenantId, name: name, active: acitve}

		tenant.deactivate()

		if tenant.active {
			t.Errorf("tenant.activa must be false, but true")
		}
	})
	t.Run("deactive to deactive", func(t *testing.T) {
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenantId := TenantId{id: uu}
		name := "TenantName"
		acitve := false
		tenant := Tenant{tenantId: tenantId, name: name, active: acitve}

		tenant.deactivate()

		if tenant.active {
			t.Errorf("tenant.activa must be false, but true")
		}
	})
}

func TestActivate(t *testing.T) {
	t.Run("deactive to active", func(t *testing.T) {
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenantId := TenantId{id: uu}
		name := "TenantName"
		acitve := false
		tenant := Tenant{tenantId: tenantId, name: name, active: acitve}

		tenant.activate()

		if !tenant.active {
			t.Errorf("tenant.activa must be true, but false")
		}
	})
	t.Run("active to active", func(t *testing.T) {
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenantId := TenantId{id: uu}
		name := "TenantName"
		acitve := true
		tenant := Tenant{tenantId: tenantId, name: name, active: acitve}

		tenant.activate()

		if !tenant.active {
			t.Errorf("tenant.activa must be true, but false")
		}
	})
}

func TestEquals(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenantId := &TenantId{id: uu}

		name1 := "TenantName1"
		tenant1 := &Tenant{tenantId: *tenantId, name: name1}

		name2 := "TenantName1"
		tenant2 := &Tenant{tenantId: *tenantId, name: name2}

		if !tenant1.Equals(*tenant2) {
			t.Errorf("tenant1 %v must be equal to %v by tenantId", tenant1, tenant2)
		}
	})
	t.Run("not equal", func(t *testing.T) {
		u1, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu1 := u1.String()
		tenantId1 := &TenantId{id: uu1}

		u2, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu2 := u2.String()
		tenantId2 := &TenantId{id: uu2}

		name := "TenantName"
		tenant1 := &Tenant{tenantId: *tenantId1, name: name}
		tenant2 := &Tenant{tenantId: *tenantId2, name: name}

		if tenant1.Equals(*tenant2) {
			t.Errorf("tenant1 %v must be not equal to %v by tenantId", tenant1, tenant2)
		}
	})
}
