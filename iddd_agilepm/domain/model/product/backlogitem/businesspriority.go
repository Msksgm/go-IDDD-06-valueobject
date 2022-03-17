package backlogitem

import (
	"fmt"
	"reflect"
)

type BusinessPriority struct {
	ratings BusinessPriorityRatings
}

func NewBusinessPriority(aBusinessPriorityRatings BusinessPriorityRatings) (*BusinessPriority, error) {
	return &BusinessPriority{ratings: aBusinessPriorityRatings}, nil
}

func (businessPriority *BusinessPriority) CostPercentage(aTotals BusinessPriorityTotals) float64 {
	return 100 * float64(businessPriority.ratings.cost) / float64(aTotals.TotalCost())
}

func (businessPriority *BusinessPriority) TotalValue(aTotals BusinessPriorityTotals) float64 {
	return float64(businessPriority.ratings.benefit) + float64(businessPriority.ratings.penalty)
}

func (businessPriority *BusinessPriority) String() string {
	return fmt.Sprintf("BusinessPriority [ratings=%v]", businessPriority.ratings)
}

func (businessPriority *BusinessPriority) Equals(other BusinessPriority) bool {
	return reflect.DeepEqual(businessPriority.ratings, other.ratings)
}
