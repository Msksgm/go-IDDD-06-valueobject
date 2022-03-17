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
	totalValue   = 200
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
	businessPriorityTotals, err := NewBusinessPriorityTotals(totalBenefit, totalCost, totalPenalty, totalRisk, totalValue)
	if err != nil {
		t.Fatal(err)
	}
	t.Run("equal", func(t *testing.T) {
		other := &BusinessPriorityTotals{totalBenefit: totalBenefit, totalCost: totalCost, totalPenalty: totalPenalty, totalRisk: totalRisk, totalValue: totalValue}

		if !businessPriorityTotals.Equals(*other) {
			t.Errorf("businessPriorityTotals %v must be equal to other %v", businessPriorityTotals, other)
		}
	})
	t.Run("not equal because totalBenefit", func(t *testing.T) {
		changedBenefit := 25
		other := &BusinessPriorityTotals{totalBenefit: changedBenefit, totalCost: totalCost, totalPenalty: totalPenalty, totalRisk: totalRisk, totalValue: totalValue}

		if businessPriorityTotals.Equals(*other) {
			t.Errorf("businessPriorityTotals %v must not be equal to other %v", businessPriorityTotals, other)
		}
	})
	t.Run("not equal because totalCost", func(t *testing.T) {
		changedCost := 25
		other := &BusinessPriorityTotals{totalBenefit: totalBenefit, totalCost: changedCost, totalPenalty: totalPenalty, totalRisk: totalRisk, totalValue: totalValue}

		if businessPriorityTotals.Equals(*other) {
			t.Errorf("businessPriorityTotals %v must not be equal to other %v", businessPriorityTotals, other)
		}
	})
	t.Run("not equal because totalPenalty", func(t *testing.T) {
		changedTotalPenalty := 25
		other := &BusinessPriorityTotals{totalBenefit: benefit, totalCost: totalCost, totalPenalty: changedTotalPenalty, totalRisk: totalRisk, totalValue: totalValue}

		if businessPriorityTotals.Equals(*other) {
			t.Errorf("businessPriorityTotals %v must not be equal to other %v", businessPriorityTotals, other)
		}
	})
	t.Run("not equal because totalRisk", func(t *testing.T) {
		changedTotalRisk := 25
		other := &BusinessPriorityTotals{totalBenefit: benefit, totalCost: totalCost, totalPenalty: totalPenalty, totalRisk: changedTotalRisk, totalValue: totalValue}

		if businessPriorityTotals.Equals(*other) {
			t.Errorf("businessPriorityTotals %v must not be equal to other %v", businessPriorityTotals, other)
		}
	})
	t.Run("not equal because totalValue", func(t *testing.T) {
		changedTotalValue := 100
		other := &BusinessPriorityTotals{totalBenefit: benefit, totalCost: totalCost, totalPenalty: totalPenalty, totalRisk: totalRisk, totalValue: changedTotalValue}

		if businessPriorityTotals.Equals(*other) {
			t.Errorf("businessPriorityTotals %v must not be equal to other %v", businessPriorityTotals, other)
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
