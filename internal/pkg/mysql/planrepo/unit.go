package planrepo

import "strings"

type Unit int

const (
	Unknown = iota +1
	Days
	Weeks
	Months
	Years
)

func (types Unit) String() string {
	return [...]string{"UNKNOWN", "DAYS", "WEEKS", "MONTHS", "YEARS"}[types-1]
}

func PlanUnitFromString(types string) Unit {
	types = strings.ToLower(types)
	switch types {
	case "days":
		return 2
	case "weeks":
		return 3
	case "months":
		return 4
	case "years":
		return 5
	default:
		return 0
	}
}