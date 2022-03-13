package entity

type Session struct {
	ID   string
	User User
}

func (s Session) IsAuthorized() bool {
	return s.User.ID != 0
}
