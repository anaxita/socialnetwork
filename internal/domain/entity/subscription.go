package entity

type Subscription struct {
	Model   Model `json:"model"`
	ModelID int64 `json:"model_id"`
	UserID  int64 `json:"user_id"`
}

type UserSubscriptions struct {
	Groups []Group `json:"groups" db:"groups"`
	Tags   []Tag   `json:"tags" db:"tags"`
	Users  []User  `json:"users" db:"users"`
}

type Model string

const (
	ModelGroup Model = "GROUP"
	ModelUser  Model = "USER"
	ModelTag   Model = "TAG"
)
