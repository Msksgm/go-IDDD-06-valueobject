package backlogitem

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewBusinessPriority(t *testing.T) {
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

func TestBusinessPriorityEquals(t *testing.T) {
	businessPriorityRatings, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
	if err != nil {
		t.Fatal(err)
	}
	businessPriority, err := NewBusinessPriority(*businessPriorityRatings)
	if err != nil {
		t.Fatal(err)
	}
	other := &BusinessPriority{ratings: *businessPriorityRatings}

	if !businessPriority.Equals(*other) {
		t.Errorf("businessPriority %v must be equal to other %v", businessPriority, other)
	}
}

func TestBusinessPriorityString(t *testing.T) {
	businessPriorityRatings, err := NewBusinessPriorityRatings(benefit, cost, penalty, risk)
	if err != nil {
		t.Fatal(err)
	}
	businessPriority, err := NewBusinessPriority(*businessPriorityRatings)
	if err != nil {
		t.Fatal(err)
	}
	got := fmt.Sprint(businessPriority)
	want := fmt.Sprintf("BusinessPriority [ratings=%v]", businessPriority.ratings)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
