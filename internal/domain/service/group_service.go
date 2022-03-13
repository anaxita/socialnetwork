// Package service prepares the data to be stored
package service

import (
	"context"
	"errors"
	"synergycommunity/internal/domain"
	"synergycommunity/internal/domain/entity"
	"time"

	"github.com/gocraft/dbr"
)

type GroupStorage interface {
	InsertGroup(ctx context.Context, e entity.Group) (int64, error)
	Groups(ctx context.Context, opts entity.Options) ([]entity.Group, int64, error)
	SelectByID(ctx context.Context, id int64) (entity.Group, error)
	UpdateGroup(ctx context.Context, e entity.Group) (entity.Group, error)
	DeleteGroup(ctx context.Context, id int64) error
	SelectUserPermissions(ctx context.Context, userID int64, groups ...int64) ([]domain.Perm, error)
	InsertUserRole(ctx context.Context, e entity.UserRole) error
	DeleteUserRole(ctx context.Context, e entity.UserRole) error
}

type GroupService struct {
	repo GroupStorage
}

func NewGroupService(group GroupStorage) *GroupService {
	return &GroupService{
		repo: group,
	}
}

// AddGroup creates a new Group.
func (s *GroupService) AddGroup(ctx context.Context, e entity.Group) (entity.Group, error) {
	e.CreatedAt = time.Now().UTC()
	e.UpdatedAt = e.CreatedAt

	id, err := s.repo.InsertGroup(ctx, e)
	if err != nil {
		return entity.Group{}, err
	}

	e.ID = id

	return e, nil
}

// ShowByID returns a Group by its ID.
func (s *GroupService) ShowByID(ctx context.Context, id int64) (entity.Group, error) {
	group, err := s.repo.SelectByID(ctx, id)
	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			return entity.Group{}, nil
		}

		return entity.Group{}, err
	}

	return group, nil
}

// Groups returns all Group objects.
func (s *GroupService) Groups(ctx context.Context, opts entity.Options) ([]entity.Group, int, error) {
	groups, count, err := s.repo.Groups(ctx, opts)
	if err != nil {
		if errors.Is(err, dbr.ErrNotFound) {
			return nil, 0, nil
		}

		return nil, 0, err
	}

	return groups, int(count), nil
}

// Edit modifies an existing Group.
func (s *GroupService) Edit(ctx context.Context, e entity.Group) (entity.Group, error) {
	e.UpdatedAt = time.Now().UTC()

	e, err := s.repo.UpdateGroup(ctx, e)
	if err != nil {
		return entity.Group{}, err
	}

	return e, nil
}

// Delete removes an existing Group.
func (s *GroupService) Delete(ctx context.Context, id int64) error {
	return s.repo.DeleteGroup(ctx, id)
}

func (s *GroupService) UserPermissions(ctx context.Context, userID int64, groups ...int64) ([]domain.Perm, error) {
	return s.repo.SelectUserPermissions(ctx, userID, groups...)
}

func (s *GroupService) SetUserRole(ctx context.Context, u entity.UserRole) error {
	return s.repo.InsertUserRole(ctx, u)
}

func (s *GroupService) UnsetUserRole(ctx context.Context, u entity.UserRole) error {
	return s.repo.DeleteUserRole(ctx, u)
}
