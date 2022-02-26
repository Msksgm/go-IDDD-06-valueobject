package enablement

import (
	"time"

	"github.com/Msksgm/go-IDDD-05-entity/iddd_common/ierrors"
)

type Enablement struct {
	enabled   bool
	startDate time.Time
	endDate   time.Time
}

func NewEnablement(aEnabled bool, aStartDate time.Time, aEndDate time.Time) (*Enablement, error) {
	enablement := new(Enablement)

	if err := ierrors.NewArgumentFalseError(aStartDate.After(aEndDate), "Enablement start and/or end date is invalid.").GetError(); err != nil {
		return nil, err
	}

	enablement.enabled = aEnabled
	enablement.startDate = aStartDate
	enablement.endDate = aEndDate
	return enablement, nil
}
