package backlogitem

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewBusinessPriorityRatings(t *testing.T) {
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
}
