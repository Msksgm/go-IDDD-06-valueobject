package identity

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewEmailAddress(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		address := "email@sample.com"
		got, err := NewEmailAddress(address)
		if err != nil {
			t.Error(err)
		}
		want := &EmailAddress{address: address}
		if diff := cmp.Diff(got, want, cmp.AllowUnexported(EmailAddress{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail the email address is required.", func(t *testing.T) {
		address := ""
		_, err := NewEmailAddress(address)
		want := fmt.Sprintf("emailaddress.NewEmailAddress(%s): The email address is required.", address)
		if got := err.Error(); want != got {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail Email address must be 100 characters or less.", func(t *testing.T) {
		address := RandString(101)
		_, err := NewEmailAddress(address)
		want := fmt.Sprintf("emailaddress.NewEmailAddress(%s): Email address must be 100 characters or less.", address)
		if got := err.Error(); want != got {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
