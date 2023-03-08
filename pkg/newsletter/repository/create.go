package repository

import (
	"context"

	"github.com/google/uuid"
)

func (r *repository) Create(
	ctx context.Context,
	userID uuid.UUID,
	blogID uuid.UUID,
	interests []String,
) (err error) {
	r.data = append(r.data, &subscriptionDBModel{
		UserID:    userID.String(),
		BlogID:    blogID.String(),
		Interests: interests,
	})

	return err
}
