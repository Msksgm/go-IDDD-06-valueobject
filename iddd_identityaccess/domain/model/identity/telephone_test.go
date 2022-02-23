package identity

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewTelephone(t *testing.T) {
	number := "09012345678"
	got, err := NewTelephone(number)
	if err != nil {
		t.Fatal(err)
	}
	want := &Telephone{number: number}

	if diff := cmp.Diff(want, got, cmp.AllowUnexported(Telephone{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
