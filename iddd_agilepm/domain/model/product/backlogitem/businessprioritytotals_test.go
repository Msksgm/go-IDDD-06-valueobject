package backlogitem

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	totalBenefit = 50
	totalCost    = 50
	totalPenalty = 50
	totalRisk    = 50
	totalValue   = 50
)

func TestNewBusinessPriorityTotals(t *testing.T) {
	got, err := NewBusinessPriorityTotals(totalBenefit, totalCost, totalPenalty, totalRisk, totalValue)
	if err != nil {
		t.Fatal(err)
	}

	want := &BusinessPriorityTotals{totalBenefit: totalBenefit, totalCost: totalCost, totalPenalty: totalPenalty, totalRisk: totalRisk, totalValue: totalValue}

	if diff := cmp.Diff(want, got, cmp.AllowUnexported(BusinessPriorityTotals{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}

func TestBusinessPriorityTotalsEquals(t *testing.T) {
	t.Run("equal", func(t *testing.T) {
		businessPriorityTotals, err := NewBusinessPriorityTotals(totalBenefit, totalCost, totalPenalty, totalRisk, totalValue)
		if err != nil {
			t.Fatal(err)
		}

		other := &BusinessPriorityTotals{totalBenefit: totalBenefit, totalCost: totalCost, totalPenalty: totalPenalty, totalRisk: totalRisk, totalValue: totalValue}

		if !businessPriorityTotals.Equals(*other) {
			t.Errorf("businessPriorityTotals %v must be equal to other %v", businessPriorityTotals, other)
		}
	})
}

func TestBusinessPriorityTotalsString(t *testing.T) {
	businessPriorityTotals, err := NewBusinessPriorityTotals(totalBenefit, totalCost, totalPenalty, totalRisk, totalValue)
	if err != nil {
		t.Fatal(err)
	}

	got := fmt.Sprint(businessPriorityTotals)
	want := fmt.Sprintf("BusinessPriorityTotals [totalBenefit=%d, totalCost=%d, totalPenalty=%d, totalRisk =%d, totalValue=%d]", totalBenefit, totalCost, totalPenalty, totalRisk, totalValue)
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
