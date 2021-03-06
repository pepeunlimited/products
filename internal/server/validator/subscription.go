package validator

import (
	"github.com/pepeunlimited/products/pkg/rpc/subscription"
	"github.com/twitchtv/twirp"
)

type SubscriptionServerValidator struct {}

func (v SubscriptionServerValidator) StartSubscription(params *subscription.StartSubscriptionParams) error {
	if params.PlanId == 0 {
		return twirp.RequiredArgumentError("plan_id")
	}
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	return nil
}

func (v SubscriptionServerValidator) StopSubscription(params *subscription.StopSubscriptionParams) error {
	return nil
}

func (v SubscriptionServerValidator) GetSubscription(params *subscription.GetSubscriptionParams) error {
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	if params.SubscriptionId == 0 {
		return twirp.RequiredArgumentError("subscription_id")
	}
	return nil
}

func (v SubscriptionServerValidator) GetSubscriptions(params *subscription.GetSubscriptionsParams) error {
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	if params.PageSize == 0 {
		return twirp.RequiredArgumentError("page_size")
	}
	return nil
}

func NewSubscriptionServerValidator() SubscriptionServerValidator {
	return SubscriptionServerValidator{}
}