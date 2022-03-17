package backlogitem

import (
	"fmt"
	"reflect"
)

type BusinessPriorityTotals struct {
	totalBenefit int
	totalCost    int
	totalPenalty int
	totalRisk    int
	totalValue   int
}

func NewBusinessPriorityTotals(aTotalBenefit int, aTotalCost int, aTotalPenalty int, aTotalRisk int, aTotalValue int) (*BusinessPriorityTotals, error) {
	return &BusinessPriorityTotals{totalBenefit: aTotalBenefit, totalCost: aTotalCost, totalPenalty: aTotalBenefit, totalRisk: aTotalRisk, totalValue: aTotalRisk}, nil
}

func (busineePriorityTotals *BusinessPriorityTotals) String() string {
	return fmt.Sprintf("BusinessPriorityTotals [totalBenefit=%d, totalCost=%d, totalPenalty=%d, totalRisk =%d, totalValue=%d]", busineePriorityTotals.totalBenefit, busineePriorityTotals.totalCost, busineePriorityTotals.totalPenalty, busineePriorityTotals.totalRisk, busineePriorityTotals.totalValue)
}

func (busineePriorityTotals *BusinessPriorityTotals) Equals(other BusinessPriorityTotals) bool {
	if !reflect.DeepEqual(busineePriorityTotals.totalBenefit, other.totalBenefit) {
		return false
	}
	if !reflect.DeepEqual(busineePriorityTotals.totalCost, other.totalCost) {
		return false
	}
	if !reflect.DeepEqual(busineePriorityTotals.totalPenalty, other.totalPenalty) {
		return false
	}
	if !reflect.DeepEqual(busineePriorityTotals.totalRisk, other.totalRisk) {
		return false
	}
	if !reflect.DeepEqual(busineePriorityTotals.totalValue, other.totalValue) {
		return false
	}
	return true
}
