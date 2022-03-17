package backlogitem

import "fmt"

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
