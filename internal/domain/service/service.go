package service

type Services struct {
	Users  *UserService
	Posts  *PostService
	Groups *GroupService
	Tags   *TagService
}

func NewServices(s Storage) *Services {
	return &Services{
		Users:  NewUserService(s),
		Posts:  NewPostService(s),
		Groups: NewGroupService(s),
		Tags:   NewTagService(s),
	}
}
