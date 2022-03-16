package identity

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/utils"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

var uuidV4 = uuid.New().String()

func TestNewTenant(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		tenantId, err := NewTenantId(uuidV4)
		if err != nil {
			t.Fatal(err)
		}

		name := "TenantName"
		got, err := NewTenant(*tenantId, name, true)
		if err != nil {
			t.Fatal(err)
		}

		want := &Tenant{tenantId: *tenantId, name: "TenantName", active: true}

		if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tenant{}, TenantId{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail empty name", func(t *testing.T) {
		tenantId, err := NewTenantId(uuidV4)
		if err != nil {
			t.Fatal(err)
		}

		name := ""
		active := true
		_, err = NewTenant(*tenantId, name, active)
		if !errors.As(err, &argumentNotEmptyError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentNotEmptyError))
		}
	})
	t.Run("fail over 100 characters name", func(t *testing.T) {
		tenantId, err := NewTenantId(uuidV4)
		if err != nil {
			t.Fatal(err)
		}

		name := utils.RandString(101)
		active := true
		_, err = NewTenant(*tenantId, name, active)
		if !errors.As(err, &argumentLengthError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentLengthError))
		}
	})
}

func TestDeactivate(t *testing.T) {
	t.Run("active to deactive", func(t *testing.T) {
		tenantId, err := NewTenantId(uuidV4)
		if err != nil {
			t.Fatal(err)
		}
		name := "TenantName"
		acitve := true
		tenant := Tenant{tenantId: *tenantId, name: name, active: acitve}

		tenant.deactivate()

		if tenant.active {
			t.Errorf("tenant.activa must be false, but true")
		}
	})
	t.Run("deactive to deactive", func(t *testing.T) {
		tenantId, err := NewTenantId(uuidV4)
		if err != nil {
			t.Fatal(err)
		}
		name := "TenantName"
		acitve := false
		tenant := Tenant{tenantId: *tenantId, name: name, active: acitve}

		tenant.deactivate()

		if tenant.active {
			t.Errorf("tenant.activa must be false, but true")
		}
	})
}

func TestActivate(t *testing.T) {
	t.Run("deactive to active", func(t *testing.T) {
		tenantId, err := NewTenantId(uuidV4)
		if err != nil {
			t.Fatal(err)
		}
		name := "TenantName"
		acitve := false
		tenant := Tenant{tenantId: *tenantId, name: name, active: acitve}

		tenant.activate()

		if !tenant.active {
			t.Errorf("tenant.activa must be true, but false")
		}
	})
	t.Run("active to active", func(t *testing.T) {
		tenantId, err := NewTenantId(uuidV4)
		if err != nil {
			t.Fatal(err)
		}
		name := "TenantName"
		acitve := true
		tenant := Tenant{tenantId: *tenantId, name: name, active: acitve}

		tenant.activate()

		if !tenant.active {
			t.Errorf("tenant.activa must be true, but false")
		}
	})
}

func TestTenantEquals(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		tenantId, err := NewTenantId(uuidV4)
		if err != nil {
			t.Fatal(err)
		}

		name1 := "TenantName1"
		tenant1 := &Tenant{tenantId: *tenantId, name: name1}

		name2 := "TenantName1"
		tenant2 := &Tenant{tenantId: *tenantId, name: name2}

		if !tenant1.Equals(*tenant2) {
			t.Errorf("tenant1 %v must be equal to %v by tenantId", tenant1, tenant2)
		}
	})
	t.Run("not equal", func(t *testing.T) {
		tenantId1, err := NewTenantId(uuidV4)
		if err != nil {
			t.Fatal(err)
		}

		otherUuidV4 := uuid.New().String()
		tenantId2, err := NewTenantId(otherUuidV4)
		if err != nil {
			t.Fatal(err)
		}

		name := "TenantName"
		tenant1 := &Tenant{tenantId: *tenantId1, name: name}
		tenant2 := &Tenant{tenantId: *tenantId2, name: name}

		if tenant1.Equals(*tenant2) {
			t.Errorf("tenant1 %v must be not equal to %v by tenantId", tenant1, tenant2)
		}
	})
}
