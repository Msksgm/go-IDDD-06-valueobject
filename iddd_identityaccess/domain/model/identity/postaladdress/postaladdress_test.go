package postaladdress

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/utils"
	"github.com/google/go-cmp/cmp"
)

var (
	argumentLengthError   *ierrors.ArgumentLengthError
	argumentNotEmptyError *ierrors.ArgumentNotEmptyError
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
		if !errors.As(err, &argumentNotEmptyError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentNotEmptyError))
		}
	})
	t.Run("fail The street address is required.", func(t *testing.T) {
		streetAddress, city, stateProvince, postalCode, countryCode := utils.RandString(101), "city", "stateProvince", "postalCode", "00"
		_, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
		if !errors.As(err, &argumentLengthError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentLengthError))
		}
	})
	t.Run("fail The street city is required.", func(t *testing.T) {
		streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", "", "stateProvince", "postalCode", "00"
		_, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
		if !errors.As(err, &argumentNotEmptyError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentNotEmptyError))
		}
	})
	t.Run("fail The street city is required.", func(t *testing.T) {
		streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", utils.RandString(101), "stateProvince", "postalCode", "00"
		_, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
		if !errors.As(err, &argumentLengthError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentLengthError))
		}
	})
	t.Run("fail The postal code is required.", func(t *testing.T) {
		streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", "city", "stateProvince", "", "00"
		_, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
		if !errors.As(err, &argumentNotEmptyError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentLengthError))
		}
	})
	t.Run("fail The postal code must be 12 characters or less.", func(t *testing.T) {
		streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", "city", "stateProvince", utils.RandString(13), "00"
		_, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
		if !errors.As(err, &argumentLengthError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentLengthError))
		}
	})
	t.Run("fail The postal code must be 5 characters or more.", func(t *testing.T) {
		streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", "city", "stateProvince", utils.RandString(4), "00"
		_, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
		if !errors.As(err, &argumentLengthError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentLengthError))
		}
	})
	t.Run("fail The state/province is required.", func(t *testing.T) {
		streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", "city", "", "postalCode", "00"
		_, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
		if !errors.As(err, &argumentNotEmptyError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentLengthError))
		}
	})
	t.Run("fail The postal code must be 12 characters or less.", func(t *testing.T) {
		streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", "city", utils.RandString(1), "postalCode", "00"
		_, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
		if !errors.As(err, &argumentLengthError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentLengthError))
		}
	})
	t.Run("fail The postal code must be 5 characters or more.", func(t *testing.T) {
		streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", "city", utils.RandString(101), "postalCode", "00"
		_, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
		if !errors.As(err, &argumentLengthError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentLengthError))
		}
	})
	t.Run("fail The country is required.", func(t *testing.T) {
		streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", "city", "stateProvince", "postalCode", ""
		_, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
		if !errors.As(err, &argumentNotEmptyError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentLengthError))
		}
	})
	t.Run("fail The country code must be two characters.", func(t *testing.T) {
		streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", "city", "stateProvince", "postalCode", "000"
		_, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
		if !errors.As(err, &argumentLengthError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentLengthError))
		}
	})
}

func TestCity(t *testing.T) {
	streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", "city", "stateProvince", "postalCode", "00"
	newPostalAddress, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
	if err != nil {
		log.Fatal(err)
	}
	got := newPostalAddress.City()
	want := city

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestCountryCode(t *testing.T) {
	streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", "city", "stateProvince", "postalCode", "00"
	newPostalAddress, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
	if err != nil {
		log.Fatal(err)
	}
	got := newPostalAddress.CountryCode()
	want := countryCode

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestPostalCode(t *testing.T) {
	streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", "city", "stateProvince", "postalCode", "00"
	newPostalAddress, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
	if err != nil {
		log.Fatal(err)
	}
	got := newPostalAddress.PostalCode()
	want := postalCode

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestStateProvince(t *testing.T) {
	streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", "city", "stateProvince", "postalCode", "00"
	newPostalAddress, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
	if err != nil {
		log.Fatal(err)
	}
	got := newPostalAddress.StateProvince()
	want := stateProvince

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestStreetAddress(t *testing.T) {
	streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", "city", "stateProvince", "postalCode", "00"
	newPostalAddress, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
	if err != nil {
		log.Fatal(err)
	}
	got := newPostalAddress.StreetAddress()
	want := streetAddress

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestPostalAddressEqual(t *testing.T) {
	streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", "city", "stateProvince", "postalCode", "00"
	newPostalAddress, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
	if err != nil {
		log.Fatal(err)
	}

	otherPostalAddress := &PostalAddress{streetAddress: streetAddress, city: city, stateProvince: stateProvince, postalCode: postalCode, countryCode: countryCode}

	if !newPostalAddress.Equals(*otherPostalAddress) {
		t.Errorf("newPostalAddress: %v must be equal to otherPostalAddress: %v", newPostalAddress, otherPostalAddress)
	}
}

func TestPostalAddressString(t *testing.T) {
	streetAddress, city, stateProvince, postalCode, countryCode := "streetAddress", "city", "stateProvince", "postalCode", "00"
	newPostalAddress, err := NewPostalAddress(streetAddress, city, stateProvince, postalCode, countryCode)
	if err != nil {
		log.Fatal(err)
	}
	got := newPostalAddress.String()
	want := fmt.Sprintf("PostalAddress [streetAddress=%s, city=%s, stateProvince=%s, postalCode=%s, countryCode=%s]", streetAddress, city, stateProvince, postalCode, countryCode)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
