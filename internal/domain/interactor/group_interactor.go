// Package interactor defines a set of logical actions which can be externally applied to the entity
package interactor

import (
	"context"
	"synergycommunity/internal/domain"
	"synergycommunity/internal/domain/entity"
	"synergycommunity/internal/domain/service"
)

type GroupInteractor struct {
	group *service.GroupService
}

func NewGroupInteractor(group *service.GroupService) *GroupInteractor {
	return &GroupInteractor{group: group}
}

// Add creates a new Group.
func (i *GroupInteractor) Add(ctx context.Context, e entity.Group) (entity.Group, error) {
	return i.group.AddGroup(ctx, e)
}

// ShowGroup return a Group object with the given ID.
func (i *GroupInteractor) ShowGroup(ctx context.Context, id int64) (entity.Group, error) {
	return i.group.ShowByID(ctx, id)
}

// ShowGroups returns all existing ShowGroups.
func (i *GroupInteractor) ShowGroups(ctx context.Context, opts entity.Options) ([]entity.Group, int, error,
) {
	return i.group.Groups(ctx, opts)
}

// Edit modifies the ShowGroups.
func (i *GroupInteractor) Edit(ctx context.Context, e entity.Group) (entity.Group, error) {
	return i.group.Edit(ctx, e)
}

// Remove removes the Group with the given ID.
func (i *GroupInteractor) Remove(ctx context.Context, id int64) error {
	return i.group.Delete(ctx, id)
}

// UserPermissions returns permissions for userID in the groups and always for the global group
func (i *GroupInteractor) UserPermissions(ctx context.Context, userID int64, groups ...int64) ([]domain.Perm, error) {
	return i.group.UserPermissions(ctx, userID, groups...)
}

// SetUserRole sets a role to a user in a group.
func (i *GroupInteractor) SetUserRole(ctx context.Context, u entity.UserRole) error {
	return i.group.SetUserRole(ctx, u)
}

// UnsetUserRole unsets a role to a user in a group.
func (i *GroupInteractor) UnsetUserRole(ctx context.Context, u entity.UserRole) error {
	return i.group.UnsetUserRole(ctx, u)
}
