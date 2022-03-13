package interactor

import "synergycommunity/internal/domain/service"

type Interactors struct {
	Users  *UserInteractor
	Posts  *PostInteractor
	Groups *GroupInteractor
	Tags   *TagInteractor
}

func NewInteractors(s *service.Services) *Interactors {
	return &Interactors{
		Users:  NewUserInteractor(s.Users),
		Posts:  NewPostInteractor(s.Posts),
		Groups: NewGroupInteractor(s.Groups),
		Tags:   NewTagInteractor(s.Tags),
	}
}
