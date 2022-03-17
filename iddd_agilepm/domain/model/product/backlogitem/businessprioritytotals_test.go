package backlogitem

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	totalBenefit = 5
	totalCost    = 5
	totalPenalty = 5
	totalRisk    = 5
	totalValue   = 5
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
