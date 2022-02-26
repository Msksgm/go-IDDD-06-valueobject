package enablement

import "time"

type Enablement struct {
	enabled   bool
	startDate time.Time
	endDate   time.Time
}

func NewEnablement(aEnabled bool, aStartDate time.Time, aEndDate time.Time) (*Enablement, error) {
	enablement := new(Enablement)

	enablement.enabled = aEnabled
	enablement.startDate = aStartDate
	enablement.endDate = aEndDate
	return enablement, nil
}
