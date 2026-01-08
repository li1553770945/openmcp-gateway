package domain

type UserEntity struct {
	ID       int64
	Username string
	Nickname string
	Role     string
	CanUse   bool
	Password string
}
