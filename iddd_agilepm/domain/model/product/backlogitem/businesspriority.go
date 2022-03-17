package backlogitem

import "reflect"

type BusinessPriority struct {
	ratings BusinessPriorityRatings
}

func NewBusinessPriority(aBusinessPriorityRatings BusinessPriorityRatings) (*BusinessPriority, error) {
	return &BusinessPriority{ratings: aBusinessPriorityRatings}, nil
}

func (businessPriority *BusinessPriority) Equals(other BusinessPriority) bool {
	return reflect.DeepEqual(businessPriority.ratings, other.ratings)
}
