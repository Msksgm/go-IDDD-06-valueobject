package identity

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func TestNewUser(t *testing.T) {
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

		userName := "userName"
		password := "qwerty!ASDFG#"
		bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
		if err != nil {
			t.Fatal(err)
		}

		got, err := NewUser(*tenantId, userName, password)
		if err != nil {
			t.Fatal(err)
		}

		want := &User{tenantId: TenantId{id: uu}, userName: userName, password: string(bcryptedPassword)}

		opts := cmp.Options{
			cmp.AllowUnexported(User{}, TenantId{}),
			cmpopts.IgnoreFields(User{}, "password"),
		}
		if diff := cmp.Diff(want, got, opts); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
		if err := bcrypt.CompareHashAndPassword([]byte(want.password), []byte(password)); err != nil {
			t.Error(err)
		}
	})
	t.Run("fail username is required.", func(t *testing.T) {
		u, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu := u.String()

		tenantId, err := NewTenantId(uu)
		if err != nil {
			t.Fatal(err)
		}

		userName := ""
		password := "qwerty!ASDFG#"

		_, err = NewUser(*tenantId, userName, password)
		want := fmt.Sprintf("user.setUserName(%s): The username is required.", userName)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestUserEquals(t *testing.T) {
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

		userName := "userName"
		password := "qwerty!ASDFG#"
		bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
		if err != nil {
			t.Fatal(err)
		}

		user, err := NewUser(*tenantId, userName, password)
		if err != nil {
			t.Fatal(err)
		}

		other := &User{tenantId: TenantId{id: uu}, userName: userName, password: string(bcryptedPassword)}

		if !user.Equals(*other) {
			t.Errorf("user: %v must be equal to other :%v", user, other)
		}
	})
	t.Run("fail tenantId is not equal", func(t *testing.T) {
		u1, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu1 := u1.String()
		tenantId1 := &TenantId{id: uu1}

		userName := "userName"
		password := "qwerty!ASDFG#"
		bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
		if err != nil {
			t.Fatal(err)
		}

		user, err := NewUser(*tenantId1, userName, password)
		if err != nil {
			t.Fatal(err)
		}

		u2, err := uuid.NewRandom()
		if err != nil {
			t.Fatal(err)
		}
		uu2 := u2.String()
		tenantId2 := &TenantId{id: uu2}

		other := &User{tenantId: *tenantId2, userName: userName, password: string(bcryptedPassword)}

		if user.Equals(*other) {
			t.Errorf("user: %v must be equal to other :%v", user, other)
		}
	})
}
