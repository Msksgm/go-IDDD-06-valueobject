package identity

import (
	"fmt"
	"testing"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/utils"
	"github.com/google/go-cmp/cmp"
)

func TestNewTelephone(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		number := "090-1234-5678"
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
		number := utils.RandString(4)
		_, err := NewTelephone(number)
		want := fmt.Sprintf("telephone.NewTelephone(%s): Telephone number must be between 5 and 20 characters.", number)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail Telephone number is more than 20 characters", func(t *testing.T) {
		number := utils.RandString(21)
		_, err := NewTelephone(number)
		want := fmt.Sprintf("telephone.NewTelephone(%s): Telephone number must be between 5 and 20 characters.", number)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
	t.Run("fail Telephone number or its format is invalid.", func(t *testing.T) {
		number := "abc-defg-heij"
		_, err := NewTelephone(number)
		want := fmt.Sprintf("telephone.NewTelephone(%s): Telephone number or its format is invalid.", number)
		if got := err.Error(); got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}

func TestTelephoneEqual(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		number := "090-1234-5678"
		telephone, err := NewTelephone(number)
		if err != nil {
			t.Fatal(err)
		}
		otherTelephone := &Telephone{number: number}

		if !telephone.Equal(otherTelephone) {
			t.Errorf("telephone: %v must be equal to otherTelephone %v", telephone, otherTelephone)
		}
	})
	t.Run("not equal", func(t *testing.T) {
		number := "090-1234-5678"
		telephone, err := NewTelephone(number)
		if err != nil {
			t.Fatal(err)
		}
		otherNumber := "090-0000-0000"
		otherTelephone := &Telephone{number: otherNumber}

		if telephone.Equal(otherTelephone) {
			t.Errorf("telephone: %v must be not equal to otherTelephone %v", telephone, otherTelephone)
		}
	})
}

func TestNumber(t *testing.T) {
	number := "090-1234-5678"
	telephone, err := NewTelephone(number)
	if err != nil {
		t.Fatal(err)
	}
	got := telephone.Number()
	want := number
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestTelephoneString(t *testing.T) {
	number := "090-1234-5678"
	telephone, err := NewTelephone(number)
	if err != nil {
		t.Fatal(err)
	}
	got := fmt.Sprint(telephone)
	want := fmt.Sprintf("Telephone [number=%v]", number)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
