package backlogitem

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
	"github.com/google/go-cmp/cmp"
)

var argumentRangeError *ierrors.ArgumentRangeError

var (
	benefit = 5
	cost    = 5
	penalty = 5
	risk    = 5
)

func TestNewBusinessPriorityRatings(t *testing.T) {
	t.Run("success", func(t *testing.T) {
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

		_, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if !errors.As(err, &argumentRangeError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentRangeError))
		}
	})
	t.Run("fail benefit is bigger than 9", func(t *testing.T) {
		benefit := 10

		_, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if !errors.As(err, &argumentRangeError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentRangeError))
		}
	})
	t.Run("fail cost is lower than 1", func(t *testing.T) {
		cost := -1

		_, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if !errors.As(err, &argumentRangeError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentRangeError))
		}
	})
	t.Run("fail cost is bigger than 9", func(t *testing.T) {
		cost := 10

		_, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if !errors.As(err, &argumentRangeError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentRangeError))
		}
	})
	t.Run("fail penalty is lower than 1", func(t *testing.T) {
		penalty := -1

		_, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if !errors.As(err, &argumentRangeError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentRangeError))
		}
	})
	t.Run("fail penalty is bigger than 9", func(t *testing.T) {
		penalty := 10

		_, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if !errors.As(err, &argumentRangeError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentRangeError))
		}
	})
	t.Run("fail risk is lower than 1", func(t *testing.T) {
		risk := -1

		_, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if !errors.As(err, &argumentRangeError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentRangeError))
		}
	})
	t.Run("fail penalty is bigger than 9", func(t *testing.T) {
		risk := 10

		_, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if !errors.As(err, &argumentRangeError) {
			t.Errorf("err type:%v, expect type: %v", reflect.TypeOf(errors.Unwrap(err)), reflect.TypeOf(&argumentRangeError))
		}
	})
}

func TestBusinessPriorityRatingsString(t *testing.T) {
	businessPriorityRatings, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
	if err != nil {
		t.Fatal(err)
	}

	got := fmt.Sprint(businessPriorityRatings)
	want := fmt.Sprintf("BusinessPriorityRatings [benefit=%d, cost=%d, penalty=%d, risk =%d]", benefit, cost, penalty, risk)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestBusinessPriorityRatingsEquals(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		businessPriorityRatings, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if err != nil {
			t.Fatal(err)
		}

		other := &BusinessPriorityRatings{benefit: benefit, cost: cost, penalty: penalty, risk: risk}

		if !businessPriorityRatings.Equals(*other) {
			t.Errorf("businessPriorityRatings %v must be equal to other %v", businessPriorityRatings, other)
		}
	})
	t.Run("not equal because benefit", func(t *testing.T) {
		businessPriorityRatings, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if err != nil {
			t.Fatal(err)
		}

		other := &BusinessPriorityRatings{benefit: 1, cost: cost, penalty: penalty, risk: risk}
		if businessPriorityRatings.Equals(*other) {
			t.Errorf("businessPriorityRatings %v must not be equal to other %v", businessPriorityRatings, other)
		}
	})
	t.Run("not equal because cost", func(t *testing.T) {
		businessPriorityRatings, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if err != nil {
			t.Fatal(err)
		}

		other := &BusinessPriorityRatings{benefit: benefit, cost: 1, penalty: penalty, risk: risk}
		if businessPriorityRatings.Equals(*other) {
			t.Errorf("businessPriorityRatings %v must not be equal to other %v", businessPriorityRatings, other)
		}
	})
	t.Run("not equal because penalty", func(t *testing.T) {
		businessPriorityRatings, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if err != nil {
			t.Fatal(err)
		}

		other := &BusinessPriorityRatings{benefit: benefit, cost: cost, penalty: 1, risk: risk}
		if businessPriorityRatings.Equals(*other) {
			t.Errorf("businessPriorityRatings %v must not be equal to other %v", businessPriorityRatings, other)
		}
	})
	t.Run("not equal because risk", func(t *testing.T) {
		businessPriorityRatings, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
		if err != nil {
			t.Fatal(err)
		}

		other := &BusinessPriorityRatings{benefit: benefit, cost: cost, penalty: penalty, risk: 1}
		if businessPriorityRatings.Equals(*other) {
			t.Errorf("businessPriorityRatings %v must not be equal to other %v", businessPriorityRatings, other)
		}
	})
}

func TestWithAdjustedBenefit(t *testing.T) {
	businessPriorityRatings, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
	if err != nil {
		t.Fatal(err)
	}

	changedBenefit := 7
	got, err := businessPriorityRatings.WithAdjustedBenefit(changedBenefit)
	if err != nil {
		t.Fatal(err)
	}
	want := &BusinessPriorityRatings{benefit: changedBenefit, cost: cost, penalty: penalty, risk: risk}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(BusinessPriorityRatings{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestWithAdjustedCost(t *testing.T) {
	businessPriorityRatings, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
	if err != nil {
		t.Fatal(err)
	}

	changedCost := 7
	got, err := businessPriorityRatings.WithAdjustedCost(changedCost)
	if err != nil {
		t.Fatal(err)
	}
	want := &BusinessPriorityRatings{benefit: benefit, cost: changedCost, penalty: penalty, risk: risk}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(BusinessPriorityRatings{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestWithAdjustedPenalty(t *testing.T) {
	businessPriorityRatings, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
	if err != nil {
		t.Fatal(err)
	}

	changedPenalty := 7
	got, err := businessPriorityRatings.WithAdjustedPenalty(changedPenalty)
	if err != nil {
		t.Fatal(err)
	}
	want := &BusinessPriorityRatings{benefit: benefit, cost: cost, penalty: changedPenalty, risk: risk}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(BusinessPriorityRatings{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestWithAdjustedRisk(t *testing.T) {
	businessPriorityRatings, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
	if err != nil {
		t.Fatal(err)
	}

	changedRisk := 7
	got, err := businessPriorityRatings.WithAdjustedRisk(changedRisk)
	if err != nil {
		t.Fatal(err)
	}
	want := &BusinessPriorityRatings{benefit: benefit, cost: cost, penalty: penalty, risk: changedRisk}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(BusinessPriorityRatings{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestBenefit(t *testing.T) {
	businessPriorityRatings, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
	if err != nil {
		t.Fatal(err)
	}
	got := businessPriorityRatings.Benefit()
	want := benefit

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCost(t *testing.T) {
	businessPriorityRatings, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
	if err != nil {
		t.Fatal(err)
	}
	got := businessPriorityRatings.Cost()
	want := benefit

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPenalty(t *testing.T) {
	businessPriorityRatings, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
	if err != nil {
		t.Fatal(err)
	}
	got := businessPriorityRatings.Penalty()
	want := benefit

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func Risk(t *testing.T) {
	businessPriorityRatings, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
	if err != nil {
		t.Fatal(err)
	}
	got := businessPriorityRatings.Risk()
	want := benefit

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
