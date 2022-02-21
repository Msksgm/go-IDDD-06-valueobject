package identity

import (
	"fmt"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewFullName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		got, err := NewFullName("FirstName", "lastName")
		if err != nil {
			log.Fatal(err)
		}

		want := &FullName{firstName: "FirstName", lastName: "lastName"}

		if diff := cmp.Diff(got, want, cmp.AllowUnexported(FullName{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail fistName empty", func(t *testing.T) {
		firstName, lastName := "", "lastName"
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): First name is required.", firstName, lastName)
		if got := err.Error(); got != want {
			log.Fatal(err)
		}
	})
	t.Run("fail fistName is over 50 characters", func(t *testing.T) {
		firstName, lastName := RandString(51), "lastName"
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): First name must be 50 characters or less.", firstName, lastName)
		if got := err.Error(); got != want {
			log.Fatal(err)
		}
	})
	t.Run("fail first Name not start with a capital letter", func(t *testing.T) {
		firstName, lastName := "firstName", "lastName"
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): First name must be at least one character in length, starting with a capital letter.", firstName, lastName)
		if got := err.Error(); got != want {
			log.Fatal(err)
		}
	})
	t.Run("fail Last name must be 50 characters or less.", func(t *testing.T) {
		firstName, lastName := "FirstName", ""
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): Last name is required.", firstName, lastName)
		if got := err.Error(); got != want {
			log.Fatal(err)
		}
	})
	t.Run("fail Last name must be 50 characters or less.", func(t *testing.T) {
		firstName, lastName := "FirstName", RandString(51)
		_, err := NewFullName(firstName, lastName)
		want := fmt.Sprintf("fullname.NewFullName(%s, %s): Last name must be 50 characters or less.", firstName, lastName)
		if got := err.Error(); got != want {
			log.Fatal(err)
		}
	})
}
