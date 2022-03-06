package identity

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/utils"
	"github.com/google/go-cmp/cmp"
)

const (
	enabled = true
)

var (
	argumentFalseError *ierrors.ArgumentFalseError
	startDate          time.Time
	endDate            time.Time
	err                error
)

func init() {
	startDate, err = time.ParseInLocation(utils.TimeFormat, "2020-01-01 00:00:00", utils.Jst)
	if err != nil {
		log.Fatal(err)
	}
	endDate, err = time.ParseInLocation(utils.TimeFormat, "2030-01-01 00:00:00", utils.Jst)
	if err != nil {
		log.Fatal(err)
	}
}

func TestNewEnablement(t *testing.T) {
	t.Run("sucess", func(t *testing.T) {
		got, err := NewEnablement(enabled, startDate, endDate)
		if err != nil {
			t.Fatal(err)
		}

		want := &Enablement{enabled: enabled, startDate: startDate, endDate: endDate}
		allowUnexported := cmp.AllowUnexported(Enablement{})
		if diff := cmp.Diff(want, got, allowUnexported); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail startDate is after than endDate", func(t *testing.T) {
		endDate, err := time.ParseInLocation(utils.TimeFormat, "2010-01-01 00:00:00", utils.Jst)
		if err != nil {
			t.Fatal(err)
		}
		_, err = NewEnablement(enabled, startDate, endDate)
		if !errors.As(err, &argumentFalseError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(err), reflect.TypeOf(&argumentFalseError))
		}
	})
}

func TestEquals(t *testing.T) {
	enablement, err := NewEnablement(enabled, startDate, endDate)
	if err != nil {
		t.Fatal(err)
	}

	otherEnablement := &Enablement{enabled: enabled, startDate: startDate, endDate: endDate}
	if !enablement.Equals(otherEnablement) {
		t.Errorf("enablement: %v must be equal to otherEnablement: %v", enablement, otherEnablement)
	}
}

func TestIsTimeExpired(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		enablement, err := NewEnablement(enabled, startDate, endDate)
		if err != nil {
			t.Fatal(err)
		}

		if enablement.IsTimeExpired() {
			t.Errorf("enablement.IsTimeExpired() must be true, when enablement: %v", enablement)
		}
	})
	t.Run("flase", func(t *testing.T) {
		startDate, err = time.ParseInLocation(utils.TimeFormat, "1900-01-01 00:00:00", utils.Jst)
		if err != nil {
			log.Fatal(err)
		}
		endDate, err = time.ParseInLocation(utils.TimeFormat, "2000-01-01 00:00:00", utils.Jst)
		if err != nil {
			log.Fatal(err)
		}
		enablement, err := NewEnablement(enabled, startDate, endDate)
		if err != nil {
			t.Fatal(err)
		}

		if !enablement.IsTimeExpired() {
			t.Errorf("enablement.IsTimeExpired() must be false, when enablement: %v", enablement)
		}
	})
}

func TestEnablementString(t *testing.T) {
	enablement, err := NewEnablement(enabled, startDate, endDate)
	if err != nil {
		t.Fatal(err)
	}
	want := fmt.Sprintf("Enablement [enabled=%v, endDate=%v, startDate=%v]", enablement.enabled, enablement.endDate, enablement.startDate)
	if got := fmt.Sprint(enablement); want != got {
		t.Errorf("got %s, want %s", got, want)
	}
}
