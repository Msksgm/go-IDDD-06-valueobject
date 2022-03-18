package backlogitem

import (
	"fmt"
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	businessPriorityRatings *BusinessPriorityRatings
	businessPriorityTotals  *BusinessPriorityTotals
	err                     error
)

func init() {
	businessPriorityRatings, err = NewBusinessPriorityRatings(benefit, cost, penalty, risk)
	if err != nil {
		log.Fatal(err)
	}
	businessPriorityTotals, err = NewBusinessPriorityTotals(totalBenefit, totalCost, totalPenalty, totalRisk, totalValue)
	if err != nil {
		log.Fatal(err)
	}
}

func TestNewBusinessPriority(t *testing.T) {
	got, err := NewBusinessPriority(*businessPriorityRatings)
	if err != nil {
		t.Fatal(err)
	}

	want := &BusinessPriority{ratings: *businessPriorityRatings}
	if diff := cmp.Diff(want, got, cmp.AllowUnexported(BusinessPriority{}, BusinessPriorityRatings{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestCostPercentage(t *testing.T) {
	businessPriority, err := NewBusinessPriority(*businessPriorityRatings)
	if err != nil {
		t.Fatal(err)
	}

	got := businessPriority.CostPercentage(*businessPriorityTotals)
	want := 100 * float64(cost) / float64(totalCost)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBusinessPriorityTotalValue(t *testing.T) {
	businessPriority, err := NewBusinessPriority(*businessPriorityRatings)
	if err != nil {
		t.Fatal(err)
	}

	got := businessPriority.TotalValue(*businessPriorityTotals)
	want := float64(benefit) + float64(penalty)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRatings(t *testing.T) {
	businessPriority, err := NewBusinessPriority(*businessPriorityRatings)
	if err != nil {
		t.Fatal(err)
	}

	got := businessPriority.Ratings()
	want := *businessPriorityRatings
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestBusinessPriorityEquals(t *testing.T) {
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
