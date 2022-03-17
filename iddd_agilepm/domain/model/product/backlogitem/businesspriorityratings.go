package backlogitem

import (
	"fmt"
	"reflect"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
)

type BusinessPriorityRatings struct {
	benefit int
	cost    int
	penalty int
	risk    int
}

func NewBusinessPriorityRatings(aBenefit int, aCost int, aPenalty int, aRisk int) (*BusinessPriorityRatings, error) {
	if err := ierrors.NewArgumentRangeError(aBenefit, 1, 9, "Relative benefit must be between 1 and 9.").GetError(); err != nil {
		return nil, err
	}
	if err := ierrors.NewArgumentRangeError(aCost, 1, 9, "Relative cost must be between 1 and 9.").GetError(); err != nil {
		return nil, err
	}
	if err := ierrors.NewArgumentRangeError(aPenalty, 1, 9, "Relative penalty must be between 1 and 9.").GetError(); err != nil {
		return nil, err
	}
	if err := ierrors.NewArgumentRangeError(aRisk, 1, 9, "Relative risk must be between 1 and 9.").GetError(); err != nil {
		return nil, err
	}

	return &BusinessPriorityRatings{benefit: aBenefit, cost: aCost, penalty: aPenalty, risk: aRisk}, nil
}

func (businessPriorityRatings *BusinessPriorityRatings) WithAdjustedBenefit(aBenefit int) (*BusinessPriorityRatings, error) {
	changedBusinessPriorityRatings, err := NewBusinessPriorityRatings(aBenefit, businessPriorityRatings.cost, businessPriorityRatings.penalty, businessPriorityRatings.risk)
	if err != nil {
		return nil, err
	}
	return changedBusinessPriorityRatings, nil
}

func (businessPriorityRatings *BusinessPriorityRatings) WithAdjustedCost(aCost int) (*BusinessPriorityRatings, error) {
	changedBusinessPriorityRatings, err := NewBusinessPriorityRatings(businessPriorityRatings.benefit, aCost, businessPriorityRatings.penalty, businessPriorityRatings.risk)
	if err != nil {
		return nil, err
	}
	return changedBusinessPriorityRatings, nil
}

func (businessPriorityRatings *BusinessPriorityRatings) WithAdjustedPenalty(aPenalty int) (*BusinessPriorityRatings, error) {
	changedBusinessPriorityRatings, err := NewBusinessPriorityRatings(businessPriorityRatings.benefit, businessPriorityRatings.cost, aPenalty, businessPriorityRatings.risk)
	if err != nil {
		return nil, err
	}
	return changedBusinessPriorityRatings, nil
}

func (businessPriorityRatings *BusinessPriorityRatings) WithAdjustedRisk(aRisk int) (*BusinessPriorityRatings, error) {
	changedBusinessPriorityRatings, err := NewBusinessPriorityRatings(businessPriorityRatings.benefit, businessPriorityRatings.cost, businessPriorityRatings.penalty, aRisk)
	if err != nil {
		return nil, err
	}
	return changedBusinessPriorityRatings, nil
}

func (businessPriorityRatings *BusinessPriorityRatings) Benefit() int {
	return businessPriorityRatings.benefit
}

func (businessPriorityRatings *BusinessPriorityRatings) Cost() int {
	return businessPriorityRatings.cost
}

func (businessPriorityRatings *BusinessPriorityRatings) Penalty() int {
	return businessPriorityRatings.penalty
}

func (businessPriorityRatings *BusinessPriorityRatings) String() string {
	return fmt.Sprintf("BusinessPriorityRatings [benefit=%d, cost=%d, penalty=%d, risk =%d]", businessPriorityRatings.benefit, businessPriorityRatings.cost, businessPriorityRatings.penalty, businessPriorityRatings.risk)
}

func (businessPriorityRatings *BusinessPriorityRatings) Equals(other BusinessPriorityRatings) bool {
	if !reflect.DeepEqual(businessPriorityRatings.benefit, other.benefit) {
		return false
	}

	if !reflect.DeepEqual(businessPriorityRatings.cost, other.cost) {
		return false
	}

	if !reflect.DeepEqual(businessPriorityRatings.penalty, other.penalty) {
		return false
	}

	if !reflect.DeepEqual(businessPriorityRatings.risk, other.risk) {
		return false
	}

	return true
}
