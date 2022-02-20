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
}

func TestUserEquals(t *testing.T) {
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
		fmt.Errorf("user: %v must be equal to other :%v", user, other)
	}
}
