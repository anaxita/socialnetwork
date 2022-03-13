// Package interactor defines a set of logical actions which can be externally applied to the entity
package interactor

import (
	"context"
	"synergycommunity/internal/domain/entity"
	"synergycommunity/internal/domain/service"
)

type UserInteractor struct {
	user *service.UserService
}

func NewUserInteractor(us *service.UserService) *UserInteractor {
	return &UserInteractor{user: us}
}

// Register creates a new User.
func (i *UserInteractor) Register(ctx context.Context, e entity.User) (entity.User, error) {
	return i.user.Add(ctx, e)
}

// UserByID return a User object with the given ID.
func (i *UserInteractor) UserByID(ctx context.Context, id int64) (entity.User, error) {
	return i.user.ByID(ctx, id)
}

// Users returns all existing Users.
func (i *UserInteractor) Users(ctx context.Context, opts entity.Options) ([]entity.User, int64,
	error) {
	return i.user.Users(ctx, opts)
}

// Edit modifies the User.
func (i *UserInteractor) Edit(ctx context.Context, e entity.User) (entity.User, error) {
	return i.user.Edit(ctx, e)
}

// Delete removes the User with the given ID.
func (i *UserInteractor) Delete(ctx context.Context, id int64) error {
	return i.user.Delete(ctx, id)
}

func (i *UserInteractor) Subscriptions(ctx context.Context, userID int64) (entity.UserSubscriptions, error) {
	return i.user.Subscriptions(ctx, userID)
}

func (i *UserInteractor) Subscribe(ctx context.Context, e entity.Subscription) error {
	return i.user.Subscribe(ctx, e)
}

func (i *UserInteractor) Unsubscribe(ctx context.Context, e entity.Subscription) error {
	return i.user.Unsubscribe(ctx, e)
}
