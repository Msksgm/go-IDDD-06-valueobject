package backlogitem

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
	"github.com/google/go-cmp/cmp"
)

var argumentRangeError *ierrors.ArgumentRangeError

func TestNewBusinessPriorityRatings(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		benefit := 5
		cost := 5
		penalty := 5
		risk := 5

		got, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if err != nil {
			t.Fatal(err)
		}

		want := &BusinessPriorityRatings{benefit: benefit, cost: cost, penalty: penalty, risk: risk}

		if diff := cmp.Diff(want, got, cmp.AllowUnexported(BusinessPriorityRatings{})); diff != "" {
			t.Errorf("mismatch (-want, +got):\n%s", diff)
		}
	})
	t.Run("fail benefit is lower than 1", func(t *testing.T) {
		benefit := -1
		cost := 5
		penalty := 5
		risk := 5

		_, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if !errors.As(err, &argumentRangeError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentRangeError))
		}
	})
	t.Run("fail benefit is bigger than 9", func(t *testing.T) {
		benefit := 10
		cost := 5
		penalty := 5
		risk := 5

		_, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if !errors.As(err, &argumentRangeError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentRangeError))
		}
	})
	t.Run("fail cost is lower than 1", func(t *testing.T) {
		benefit := 5
		cost := -1
		penalty := 5
		risk := 5

		_, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if !errors.As(err, &argumentRangeError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentRangeError))
		}
	})
	t.Run("fail cost is bigger than 9", func(t *testing.T) {
		benefit := 5
		cost := 10
		penalty := 5
		risk := 5

		_, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if !errors.As(err, &argumentRangeError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentRangeError))
		}
	})
	t.Run("fail penalty is lower than 1", func(t *testing.T) {
		benefit := 5
		cost := 5
		penalty := -1
		risk := 5

		_, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if !errors.As(err, &argumentRangeError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentRangeError))
		}
	})
	t.Run("fail penalty is bigger than 9", func(t *testing.T) {
		benefit := 5
		cost := 5
		penalty := 10
		risk := 5

		_, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if !errors.As(err, &argumentRangeError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentRangeError))
		}
	})
	t.Run("fail risk is lower than 1", func(t *testing.T) {
		benefit := 5
		cost := 5
		penalty := 5
		risk := -1

		_, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if !errors.As(err, &argumentRangeError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentRangeError))
		}
	})
	t.Run("fail penalty is bigger than 9", func(t *testing.T) {
		benefit := 5
		cost := 5
		penalty := 5
		risk := 10

		_, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if !errors.As(err, &argumentRangeError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentRangeError))
		}
	})
}
