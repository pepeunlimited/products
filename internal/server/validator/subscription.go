package validator

import (
	"github.com/pepeunlimited/prices/pkg/subscriptionrpc"
	"github.com/twitchtv/twirp"
)

type SubscriptionServerValidator struct {}

func (v SubscriptionServerValidator) StartSubscription(params *subscriptionrpc.StartSubscriptionParams) error {
	if params.PlanId == 0 {
		return twirp.RequiredArgumentError("plan_id")
	}
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	return nil
}

func (v SubscriptionServerValidator) StopSubscription(params *subscriptionrpc.StopSubscriptionParams) error {
	return nil
}

func (v SubscriptionServerValidator) GetSubscription(params *subscriptionrpc.GetSubscriptionParams) error {
	if params.UserId == 0 {
		return twirp.RequiredArgumentError("user_id")
	}
	if params.SubscriptionId == 0 {
		return twirp.RequiredArgumentError("subscription_id")
	}
	return nil
}

func (v SubscriptionServerValidator) GetSubscriptions(params *subscriptionrpc.GetSubscriptionsParams) error {
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