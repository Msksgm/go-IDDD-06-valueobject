package enablement

import (
	"errors"
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
	"github.com/google/go-cmp/cmp"
)

const (
	timeFormat = "2006-01-02 15:04:05"
	enabled    = true
)

var (
	jst                = time.FixedZone("Asia/Tokyo", 9*60*60)
	argumentFalseError *ierrors.ArgumentFalseError
	startDate          time.Time
	endDate            time.Time
	err                error
)

func init() {
	startDate, err = time.ParseInLocation(timeFormat, "2020-01-01 00:00:00", jst)
	if err != nil {
		log.Fatal(err)
	}
	endDate, err = time.ParseInLocation(timeFormat, "2030-01-01 00:00:00", jst)
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
		endDate, err := time.ParseInLocation(timeFormat, "2010-01-01 00:00:00", jst)
		if err != nil {
			t.Fatal(err)
		}
		_, err = NewEnablement(enabled, startDate, endDate)
		if !errors.As(err, &argumentFalseError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(err), reflect.TypeOf(&argumentFalseError))
		}
	})
}
