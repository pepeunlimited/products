package planrepo

import (
	"testing"
)

func TestPlanUnitFromString(t *testing.T) {
	units := PlanUnitFromString("weeks")
	if units != Weeks {
		t.Error(units)
		t.FailNow()
	}
}