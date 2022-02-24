package emailaddress

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/utils"
	"github.com/google/go-cmp/cmp"
)

var (
	argumentLengthError   *ierrors.ArgumentLengthError
	argumentNotEmptyError *ierrors.ArgumentNotEmptyError
	argumentTrueError     *ierrors.ArgumentTrueError
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
		if !errors.As(err, &argumentNotEmptyError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentNotEmptyError))
		}
	})
	t.Run("fail Email address must be 100 characters or less.", func(t *testing.T) {
		address := utils.RandString(101)
		_, err := NewEmailAddress(address)
		if !errors.As(err, &argumentLengthError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentLengthError))
		}
	})
	t.Run("fail Email address format is invalid.", func(t *testing.T) {
		address := "badmail"
		_, err := NewEmailAddress(address)
		if !errors.As(err, &argumentTrueError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentTrueError))
		}
	})
}

func TestAddress(t *testing.T) {
	address := "email@sample.com"
	newAddress, err := NewEmailAddress(address)
	if err != nil {
		t.Error(err)
	}
	got := newAddress.Address()
	want := address
	if want != got {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestEmailAddressEquals(t *testing.T) {
	address := "email@sample.com"
	newEmailAddress, err := NewEmailAddress(address)
	if err != nil {
		t.Error(err)
	}
	otherEmailAddress := &EmailAddress{address: address}
	if !newEmailAddress.Equals(*otherEmailAddress) {
		t.Errorf("newEmailAddress: %v must be equal to %v", newEmailAddress, otherEmailAddress)
	}
}

func TestEmailAddressString(t *testing.T) {
	address := "email@sample.com"
	newEmailAddress, err := NewEmailAddress(address)
	if err != nil {
		t.Error(err)
	}
	got := fmt.Sprint(newEmailAddress)
	want := fmt.Sprintf("EmailAddress [address=%s]", address)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
