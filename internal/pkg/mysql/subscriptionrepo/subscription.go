package subscriptionrepo

import (
	"context"
	"errors"
	"github.com/pepeunlimited/prices/internal/pkg/ent"
	"github.com/pepeunlimited/prices/internal/pkg/ent/subscription"
	"time"
)

var (
	ErrSubscriptionNotExist = errors.New("subscriptions: not exist")
)

type SubscriptionRepository interface {
	Create(ctx context.Context, userID int64, startAt time.Time, endAt time.Time, planID int) (*ent.Subscription, error)

	GetSubscriptionByID(ctx context.Context, subscriptionID int) (*ent.Subscription, error)
	GetSubscriptionPlanByID(ctx context.Context, subscriptionID int) (*ent.Subscription, *ent.Plan, error)
	GetSubscriptionByUserID(ctx context.Context, userID int64, pageToken int64, pageSize int32) ([]*ent.Subscription, int64, error)
}

type subscriptionMySQL struct {
	client *ent.Client
}

func (mysql subscriptionMySQL) GetSubscriptionByID(ctx context.Context, subscriptionID int) (*ent.Subscription, error) {
	selected, err := mysql.client.Subscription.Get(ctx, subscriptionID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrSubscriptionNotExist
		}
		return nil, err
	}
	return selected, nil
}

func (mysql subscriptionMySQL) GetSubscriptionPlanByID(ctx context.Context, subscriptionID int) (*ent.Subscription, *ent.Plan, error) {
	sub, err := mysql.GetSubscriptionByID(ctx, subscriptionID)
	if err != nil {
		return nil, nil, err
	}
	p, err := sub.Edges.PlansOrErr()
	if err != nil {
		return nil, nil, err
	}
	return sub, p, nil
}

func (mysql subscriptionMySQL) GetSubscriptionByUserID(ctx context.Context, userID int64, pageToken int64, pageSize int32) ([]*ent.Subscription, int64, error) {
	subs, err := mysql.client.Subscription.Query().Where(
		subscription.UserID(userID),
		subscription.IDGT(int(pageToken))).
		Order(ent.Asc(subscription.FieldID)).
		Limit(int(pageSize)).All(ctx)
	if err != nil {
		return nil, 0, err
	}
	if len(subs) == 0 {
		return []*ent.Subscription{}, pageToken, nil
	}
	return subs, int64(subs[len(subs) - 1].ID), nil
}

func (mysql subscriptionMySQL) Create(ctx context.Context, userID int64, startAt time.Time, endAt time.Time, planID int) (*ent.Subscription, error) {
	save, err := mysql.
		client.
		Subscription.
		Create().
		SetUserID(userID).
		SetEndAt(endAt.UTC()).
		SetStartAt(startAt.UTC()).
		SetPlansID(planID).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return save, nil
}

func NewSubscriptionRepository(client *ent.Client) SubscriptionRepository {
	return subscriptionMySQL{client:client}
}