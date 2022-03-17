package backlogitem

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewBusinessPriority(t *testing.T) {
	benefit := 10
	cost := 10
	penalty := 10
	risk := 10

	businessPriorityRatings, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
	if err != nil {
		t.Fatal(err)
	}
	got, err := NewBusinessPriority(*businessPriorityRatings)
	if err != nil {
		t.Fatal(err)
	}

	want := &BusinessPriority{ratings: *businessPriorityRatings}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(BusinessPriority{}, BusinessPriorityRatings{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
