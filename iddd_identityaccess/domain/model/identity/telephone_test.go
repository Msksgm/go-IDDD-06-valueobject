package identity

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewTelephone(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		number := "09012345678"
		got, err := NewTelephone(number)
		if err != nil {
			t.Fatal(err)
		}
		want := &Telephone{number: number}

		if diff := cmp.Diff(want, got, cmp.AllowUnexported(Telephone{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail Telephone number is required.", func(t *testing.T) {
		number := ""
		_, err := NewTelephone(number)
		want := fmt.Sprintf("telephone.NewTelephone(%s): Telephone number is required.", number)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail Telephone number is less than 5 characters.", func(t *testing.T) {
		number := RandString(4)
		_, err := NewTelephone(number)
		want := fmt.Sprintf("telephone.NewTelephone(%s): Telephone number must be between 5 and 20 characters.", number)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail Telephone number is more than 20 characters", func(t *testing.T) {
		number := RandString(21)
		_, err := NewTelephone(number)
		want := fmt.Sprintf("telephone.NewTelephone(%s): Telephone number must be between 5 and 20 characters.", number)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
