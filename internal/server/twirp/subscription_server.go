package twirp

import (
	"context"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/plan"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/price"
	"github.com/pepeunlimited/prices/internal/pkg/mysql/subscription"
	"github.com/pepeunlimited/prices/internal/server/errorz"
	"github.com/pepeunlimited/prices/internal/server/validator"
	"github.com/pepeunlimited/prices/pkg/rpc/subscription"
	"time"
)

type SubscriptionServer struct {
	valid         validator.SubscriptionServerValidator
	plans         plan.PlanRepository
	subscriptions subscription.SubscriptionRepository
	prices        price.PriceRepository
}

func (server SubscriptionServer) StartSubscription(ctx context.Context, params *subscription.StartSubscriptionParams) (*subscription.Subscription, error) {
	err := server.valid.StartSubscription(params)
	if err != nil {
		return nil, err
	}
	_, err = server.prices.GetPriceByPlanId(ctx, int(params.PlanId), false, false, false)
	if err != nil {
		return nil, errorz.Price(err)
	}
	now := time.Now().UTC()
	endAt, err := server.plans.EndAtByPlanID(ctx, now, int(params.PlanId))
	if err != nil {
		return nil, errorz.Plan(err)
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

func NewSubscriptionServer(client *ent.Client) SubscriptionServer {
	return SubscriptionServer{
		valid:         validator.NewSubscriptionServerValidator(),
		plans:         plan.NewPlanRepository(client),
		subscriptions: subscription.New(client),
		prices:        price.New(client),
	}
}