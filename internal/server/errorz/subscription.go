package errorz

import (
	"github.com/pepeunlimited/prices/internal/pkg/mysql/subscription"
	"github.com/twitchtv/twirp"
	"log"
)

func Subscription(err error) error {
	switch err {
	case subscription.ErrSubscriptionNotExist:
		return twirp.NotFoundError("subscription_not_found")
	}
	log.Print("subscription-service: unknown error: "+err.Error())
	return twirp.InternalErrorWith(err)
}