package identity

import (
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewPostalAddress(t *testing.T) {
	streetAddress, city, stateProvince, postalCode, countryCode := "", "", "", "", ""
	got, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
	if err != nil {
		log.Fatal(err)
	}
	want := &PostalAddress{streetAddress: streetAddress, city: city, stateProvince: stateProvince, postalCode: postalCode, countryCode: countryCode}

	if diff := cmp.Diff(want, got, cmp.AllowUnexported(PostalAddress{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
