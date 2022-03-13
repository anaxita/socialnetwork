package domain

// Perm - is a permission for access some action.
type Perm int64

const (
	PermNoReading    Perm = iota + 1 // 1 banned from reading
	PermNoWriting                    // 2 banned from writing
	PermEdit                         // 3 can edit posts and groups that don't belong to this user (in one group)
	PermDelete                       // 4 can delete an entity
	PermAdministrate                 // 5 can do everything with an entity
	PermWrite                        // 6 can write posts, comments and etc
)

type Role int64

const (
	RoleRadonly Role = iota + 1 // 1
	RoleBanned                  // 2
	RoleAdmins                  // 3
	RoleEditors                 // 4
	RoleMembers                 // 5
)

// Constants for filter conditions.
const (
	Limit    = 20
	Page     = 1
	LimitMin = 1

	ASC      = "asc"
	DESC     = "desc"
	SignEq   = "="
	SignGt   = ">"
	SignLt   = "<"
	SignLike = "like"
)
