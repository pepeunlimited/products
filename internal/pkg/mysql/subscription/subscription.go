package subscription

import (
	"context"
	"errors"
	"github.com/pepeunlimited/products/internal/pkg/ent"
	"github.com/pepeunlimited/products/internal/pkg/ent/subscription"
	"time"
)

var (
	ErrSubscriptionNotExist = errors.New("subscriptions: not exist")
)

type SubscriptionRepository interface {
	Create(ctx context.Context, userID int64, startAt time.Time, endAt time.Time, planID int) (*ent.Subscription, error)

	GetSubscriptionByIdAndUserId(ctx context.Context, subscriptionId int, userId int64, withPlans bool) (*ent.Subscription, error)
	GetSubscriptionByID(ctx context.Context, subscriptionID int, withPlans bool) (*ent.Subscription, error)
	GetSubscriptionByUserID(ctx context.Context, userID int64, pageToken int64, pageSize int32) ([]*ent.Subscription, int64, error)

	Wipe(ctx context.Context)
}

type subscriptionMySQL struct {
	client *ent.Client
}

func (mysql subscriptionMySQL) GetSubscriptionByIdAndUserId(ctx context.Context, subscriptionId int, userId int64, withPlans bool) (*ent.Subscription, error) {
	query := mysql.client.Subscription.Query().Where(subscription.ID(subscriptionId), subscription.UserID(userId))
	if withPlans {
		query.WithPlans()
	}
	selected, err := query.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrSubscriptionNotExist
		}
		return nil, err
	}
	return selected, nil
}

func (mysql subscriptionMySQL) Wipe(ctx context.Context) {
	mysql.client.Subscription.Delete().ExecX(ctx)
	mysql.client.Price.Delete().ExecX(ctx)
	mysql.client.Plan.Delete().ExecX(ctx)
	mysql.client.ThirdPartyPrice.Delete().ExecX(ctx)
	mysql.client.Product.Delete().ExecX(ctx)
}

func (mysql subscriptionMySQL) GetSubscriptionByID(ctx context.Context, subscriptionID int, withPlans bool) (*ent.Subscription, error) {
	query := mysql.client.Subscription.Query().Where(subscription.ID(subscriptionID))
	if withPlans {
		query.WithPlans()
	}
	selected, err := query.Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, ErrSubscriptionNotExist
		}
		return nil, err
	}
	return selected, nil
}

func (mysql subscriptionMySQL) GetSubscriptionByUserID(ctx context.Context, userID int64, pageToken int64, pageSize int32) ([]*ent.Subscription, int64, error) {
	subs, err := mysql.client.Subscription.Query().WithPlans().Where(
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
	return mysql.GetSubscriptionByID(ctx, save.ID, true)
}

func New(client *ent.Client) SubscriptionRepository {
	return subscriptionMySQL{client:client}
}