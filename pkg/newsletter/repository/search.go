package repository

import (
	"context"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	uuid "github.com/google/uuid"
)

func (r *repository) Search(
	ctx context.Context,
	userID uuid.UUID,
	blogID uuid.UUID,
	interests []newsletter.Interest,
	limit int,
	offset int,
) ([]*newsletter.Subscription, error) {
	var info []*newsletter.Subscription

	for _, subscription := range r.data {
		if subscription.UserID == userID.String() {
			info = append(info, &newsletter.Subscription{
				UserID:    userID,
				BlogID:    blogID,
				Interests: interests,
			})
		}
	}

	return info, nil
}
