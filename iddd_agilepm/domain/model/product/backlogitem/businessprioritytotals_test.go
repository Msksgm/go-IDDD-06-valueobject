package backlogitem

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewBusinessPriorityTotals(t *testing.T) {
	totalBenefit, totalCost, totalPenalty, totalRisk, totalValue := 5, 5, 5, 5, 5

	got, err := NewBusinessPriorityTotals(totalBenefit, totalCost, totalPenalty, totalRisk, totalValue)
	if err != nil {
		t.Fatal(err)
	}

	want := &BusinessPriorityTotals{totalBenefit: totalBenefit, totalCost: totalCost, totalPenalty: totalPenalty, totalRisk: totalRisk, totalValue: totalValue}

	if diff := cmp.Diff(want, got, cmp.AllowUnexported(BusinessPriorityTotals{})); diff != "" {
		t.Errorf("mismatch (-want, +got):\n%s", diff)
	}
}
