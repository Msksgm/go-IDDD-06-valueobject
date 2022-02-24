package tenant

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/utils"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_identityaccess/domain/model/identity/tenantid"
	"github.com/google/go-cmp/cmp"
	"github.com/google/uuid"
)

var (
	argumentLengthError   *ierrors.ArgumentLengthError
	argumentNotEmptyError *ierrors.ArgumentNotEmptyError
)

func TestNewTenant(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenantId, err := tenantid.NewTenantId(uu)
		if err != nil {
			t.Fatal(err)
		}

		name := "TenantName"
		got, err := NewTenant(*tenantId, name, true)
		if err != nil {
			t.Fatal(err)
		}

		want := &Tenant{tenantId: *tenantId, name: "TenantName", active: true}

		if diff := cmp.Diff(want, got, cmp.AllowUnexported(Tenant{}, tenantid.TenantId{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail empty name", func(t *testing.T) {
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenantId, err := tenantid.NewTenantId(uu)
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
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenantId, err := tenantid.NewTenantId(uu)
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
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenantId, err := tenantid.NewTenantId(uu)
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
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenantId, err := tenantid.NewTenantId(uu)
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
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenantId, err := tenantid.NewTenantId(uu)
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
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenantId, err := tenantid.NewTenantId(uu)
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

func TestEquals(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenantId, err := tenantid.NewTenantId(uu)
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
		u1, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu1 := u1.String()
		tenantId1, err := tenantid.NewTenantId(uu1)
		if err != nil {
			t.Fatal(err)
		}

		u2, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu2 := u2.String()
		tenantId2, err := tenantid.NewTenantId(uu2)
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
