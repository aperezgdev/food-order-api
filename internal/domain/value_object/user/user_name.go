package user_vo

type UserName string

func NewUserName(name string) UserName {
	return UserName(name)
}

func (u *UserName) Validate() bool {
	return len(*u) > 2
}
