package backlogitem

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
