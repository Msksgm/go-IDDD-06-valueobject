package enablement

import (
	"fmt"
	"reflect"
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

func (enablement *Enablement) IsEnabled() bool {
	return enablement.enabled
}

func (enablement *Enablement) EndDate() time.Time {
	return enablement.endDate
}

func (enablement *Enablement) StartDate() time.Time {
	return enablement.startDate
}

func (enablement *Enablement) IsTimeExpired() bool {
	timeExpired := false

	now := time.Now()
	if now.Before(enablement.startDate) || now.After(enablement.endDate) {
		timeExpired = true
	}

	return timeExpired
}

func (enablement *Enablement) Equals(otheEnablement *Enablement) bool {
	isEnabledEqual := reflect.DeepEqual(enablement.enabled, otheEnablement.enabled)
	isStartDateEqual := reflect.DeepEqual(enablement.startDate, enablement.startDate)
	isEndDateEqual := reflect.DeepEqual(enablement.endDate, enablement.endDate)
	return isEnabledEqual && isStartDateEqual && isEndDateEqual
}

func (enablement *Enablement) String() string {
	return fmt.Sprintf("Enablement [enabled=%v, endDate=%v, startDate=%v]", enablement.enabled, enablement.endDate, enablement.startDate)
}
