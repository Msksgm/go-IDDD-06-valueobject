package identity

import (
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewFullName(t *testing.T) {
	got, err := NewFullName("FirstName", "lastName")
	if err != nil {
		log.Fatal(err)
	}

	want := &FullName{firstName: "FirstName", lastName: "lastName"}

	if diff := cmp.Diff(got, want, cmp.AllowUnexported(FullName{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
