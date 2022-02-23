package identity

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewEmailAddress(t *testing.T) {
	address := "email@sample.com"
	got, err := NewEmailAddress(address)
	if err != nil {
		t.Error(err)
	}
	want := &EmailAddress{address: address}
	if diff := cmp.Diff(got, want, cmp.AllowUnexported(EmailAddress{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
