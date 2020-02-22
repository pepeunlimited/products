package twirp

import (
	"context"
	"github.com/pepeunlimited/products/internal/pkg/ent"
	"github.com/pepeunlimited/products/internal/pkg/mysql/plan"
	"github.com/pepeunlimited/products/internal/pkg/mysql/price"
	subscriptionrepo "github.com/pepeunlimited/products/internal/pkg/mysql/subscription"
	"github.com/pepeunlimited/products/internal/server/errorz"
	"github.com/pepeunlimited/products/internal/server/validator"
	"github.com/pepeunlimited/products/pkg/rpc/subscription"
	"github.com/twitchtv/twirp"
	"time"
)

type SubscriptionServer struct {
	valid         validator.SubscriptionServerValidator
	plans         plan.PlanRepository
	subscriptions subscriptionrepo.SubscriptionRepository
	prices        price.PriceRepository
}

func (server SubscriptionServer) StartSubscription(ctx context.Context, params *subscription.StartSubscriptionParams) (*subscription.Subscription, error) {
	err := server.valid.StartSubscription(params)
	if err != nil {
		return nil, err
	}
	plans, err := server.plans.GetPlanByID(ctx, int(params.PlanId))
	if err != nil {
		return nil, errorz.Plan(err)
	}
	now := time.Now().UTC()
	endAt, err := server.endAt(int(plans.Length), now, plan.PlanUnitFromString(plans.Unit))
	if err != nil {
		return nil, err
	}
	subscription, err := server.subscriptions.Create(ctx, params.UserId, now, endAt, int(params.PlanId))
	if err != nil {
		return nil, errorz.Subscription(err)
	}
	return ToSubscription(subscription), nil
}

func (server SubscriptionServer) StopSubscription(ctx context.Context, params *subscription.StopSubscriptionParams) (*subscription.Subscription, error) {
	err := server.valid.StopSubscription(params)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (server SubscriptionServer) GetSubscription(ctx context.Context, params *subscription.GetSubscriptionParams) (*subscription.Subscription, error) {
	err := server.valid.GetSubscription(params)
	if err != nil {
		return nil, err
	}
	subscription, err := server.subscriptions.GetSubscriptionByIdAndUserId(ctx, int(params.SubscriptionId), params.UserId, true)
	if err != nil {
		return nil, errorz.Subscription(err)
	}
	return ToSubscription(subscription), nil
}

func (server SubscriptionServer) GetSubscriptions(ctx context.Context, params *subscription.GetSubscriptionsParams) (*subscription.GetSubscriptionsResponse, error) {
	err := server.valid.GetSubscriptions(params)
	if err != nil {
		return nil, err
	}
	subscriptions, nextPageToken, err := server.subscriptions.GetSubscriptionByUserID(ctx, params.UserId, params.PageToken, params.PageSize)
	if err != nil {
		return nil, errorz.Subscription(err)
	}
	return &subscription.GetSubscriptionsResponse{
		Subscriptions: ToSubscriptions(subscriptions),
		NextPageToken: nextPageToken,
	}, nil
}

func (server SubscriptionServer) endAt(length int, startAt time.Time, unit plan.Unit) (time.Time, error) {
	switch unit {
	case plan.Days:
		return startAt.AddDate(0, 0, length).UTC(), nil
	case plan.Weeks:
		return startAt.AddDate(0, 0, length * 7).UTC(), nil
	case plan.Months:
		return startAt.AddDate(0, length, 0).UTC(),nil
	case plan.Years:
		return startAt.AddDate(length, 0, 0).UTC(),nil
	default:
		return time.Time{}, twirp.InvalidArgumentError("unit", "unknown")
	}
}

func NewSubscriptionServer(client *ent.Client) SubscriptionServer {
	return SubscriptionServer{
		valid:         validator.NewSubscriptionServerValidator(),
		plans:         plan.New(client),
		subscriptions: subscriptionrepo.New(client),
		prices:        price.New(client),
	}
}