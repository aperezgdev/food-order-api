package user_vo

import (
	"regexp"
)

type UserEmail string

var email_regex = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"

func NewUserEmail(email string) UserEmail {
	return UserEmail(email)
}

func (e *UserEmail) Validate() bool {
	pattern := regexp.MustCompile(email_regex)

	return pattern.MatchString(string(*e))
}
