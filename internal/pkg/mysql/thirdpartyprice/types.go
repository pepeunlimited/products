package thirdpartyprice

import "strings"

type Types int

const (
	Unknown = iota +1
	Consumable
	NonConsumable
	AutoRenewableSubscriptions
	NonRenewingSubscriptions
)

func (types Types) String() string {
	return [...]string{"UNKNOWN", "CONSUMABLE", "NON_CONSUMABLE", "AUTO_RENEWABLE_SUBSCRIPTIONS", "NON_RENEWING_SUBSCRIPTIONS"}[types-1]
}

func ThirdPartyPriceTypeFromString(types string) Types {
	types = strings.ToLower(types)
	switch types {
	case "consumable":
		return 2
	case "non_consumable":
		return 3
	case "auto_renewable_subscriptions":
		return 4
	case "non_renewing_subscriptions":
		return 5
	default:
		return 1
	}
}