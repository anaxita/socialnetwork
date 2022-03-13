// Package service prepares the data to be stored
package service

import (
	"context"
	"errors"
	"synergycommunity/internal/domain"
	"synergycommunity/internal/domain/entity"
	"time"
)

type UserStorage interface {
	InsertUser(ctx context.Context, e entity.User) (int64, error)
	Users(ctx context.Context, opts entity.Options) ([]entity.User, int64, error)
	UserByID(ctx context.Context, id int64) (entity.User, error)
	UpdateUser(ctx context.Context, e entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int64) error
	InsertSubscription(ctx context.Context, e entity.Subscription) error
	DeleteSubscription(ctx context.Context, e entity.Subscription) error
	Subscriptions(ctx context.Context, userID int64) (entity.UserSubscriptions, error)
}

type UserService struct {
	repo UserStorage
}

func NewUserService(repo UserStorage) *UserService {
	return &UserService{repo: repo}
}

// Add creates a new User.
func (s *UserService) Add(ctx context.Context, e entity.User) (entity.User, error) {
	e.CreatedAt = time.Now().UTC()
	e.UpdatedAt = e.CreatedAt

	id, err := s.repo.InsertUser(ctx, e)
	if err != nil {
		return entity.User{}, err
	}

	e.ID = id

	return e, nil
}

// ByID returns a User by its ID.
func (s *UserService) ByID(ctx context.Context, id int64) (entity.User, error) {
	user, err := s.repo.UserByID(ctx, id)
	if err != nil {
		return entity.User{}, err
	}

	if user.ID == 0 {
		return entity.User{}, domain.NewErrorWrap(err, domain.ErrCodeNotFound, "User", "ID", id)
	}

	return user, nil
}

// Users returns all User objects.
func (s *UserService) Users(ctx context.Context, opts entity.Options) ([]entity.User, int64, error) {
	users, count, err := s.repo.Users(ctx, opts)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, 0, nil
		}

		return nil, 0, err
	}

	return users, count, nil
}

// Edit modifies an existing User.
func (s *UserService) Edit(ctx context.Context, e entity.User) (entity.User, error) {
	e.UpdatedAt = time.Now().UTC()

	return s.repo.UpdateUser(ctx, e)
}

// Delete removes an existing User.
func (s *UserService) Delete(ctx context.Context, id int64) error {
	return s.repo.DeleteUser(ctx, id)
}

func (s *UserService) Subscriptions(ctx context.Context, userID int64) (entity.UserSubscriptions, error) {
	return s.repo.Subscriptions(ctx, userID)
}

func (s *UserService) Subscribe(ctx context.Context, e entity.Subscription) (err error) {
	return s.repo.InsertSubscription(ctx, e)
}

func (s *UserService) Unsubscribe(ctx context.Context, e entity.Subscription) (err error) {
	return s.repo.DeleteSubscription(ctx, e)

}
