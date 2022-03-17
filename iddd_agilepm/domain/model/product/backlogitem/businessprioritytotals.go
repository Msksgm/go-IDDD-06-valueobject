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
	return &BusinessPriorityTotals{totalBenefit: aTotalBenefit, totalCost: aTotalCost, totalPenalty: aTotalPenalty, totalRisk: aTotalRisk, totalValue: aTotalValue}, nil
}

func (busineePriorityTotals *BusinessPriorityTotals) TotalBenefit() int {
	return busineePriorityTotals.totalBenefit
}

func (busineePriorityTotals *BusinessPriorityTotals) TotalCost() int {
	return busineePriorityTotals.totalCost
}

func (busineePriorityTotals *BusinessPriorityTotals) TotalPenalty() int {
	return busineePriorityTotals.totalPenalty
}

func (busineePriorityTotals *BusinessPriorityTotals) TotalRisk() int {
	return busineePriorityTotals.totalRisk
}

func (busineePriorityTotals *BusinessPriorityTotals) TotalValue() int {
	return busineePriorityTotals.totalValue
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

func (busineePriorityTotals *BusinessPriorityTotals) String() string {
	return fmt.Sprintf("BusinessPriorityTotals [totalBenefit=%d, totalCost=%d, totalPenalty=%d, totalRisk =%d, totalValue=%d]", busineePriorityTotals.totalBenefit, busineePriorityTotals.totalCost, busineePriorityTotals.totalPenalty, busineePriorityTotals.totalRisk, busineePriorityTotals.totalValue)
}
