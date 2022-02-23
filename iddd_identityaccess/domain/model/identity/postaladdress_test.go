package identity

import (
	"fmt"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewPostalAddress(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", "city", "stateProvince", "postalCode", "00"
		got, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
		if err != nil {
			log.Fatal(err)
		}
		want := &PostalAddress{streetAddress: streetAddress, city: city, stateProvince: stateProvince, postalCode: postalCode, countryCode: countryCode}

		if diff := cmp.Diff(want, got, cmp.AllowUnexported(PostalAddress{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail The street address is required.", func(t *testing.T) {
		streetAddress, city, stateProvince, postalCode, countryCode := "", "city", "stateProvince", "postalCode", "00"
		_, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
		want := fmt.Sprintf("postaladdress.NewPostalAddress(%s, %s, %s, %s, %s): The street address is required.", streetAddress, city, stateProvince, postalCode, countryCode)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail The street address is required.", func(t *testing.T) {
		streetAddress, city, stateProvince, postalCode, countryCode := RandString(101), "city", "stateProvince", "postalCode", "00"
		_, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
		want := fmt.Sprintf("postaladdress.NewPostalAddress(%s, %s, %s, %s, %s): The street address must be 100 characters or less.", streetAddress, city, stateProvince, postalCode, countryCode)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail The street city is required.", func(t *testing.T) {
		streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", "", "stateProvince", "postalCode", "00"
		_, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
		want := fmt.Sprintf("postaladdress.NewPostalAddress(%s, %s, %s, %s, %s): The street city is required.", streetAddress, city, stateProvince, postalCode, countryCode)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail The street city is required.", func(t *testing.T) {
		streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", RandString(101), "stateProvince", "postalCode", "00"
		_, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
		want := fmt.Sprintf("postaladdress.NewPostalAddress(%s, %s, %s, %s, %s): The street city must be 100 characters or less.", streetAddress, city, stateProvince, postalCode, countryCode)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
