package users

import "time"

type UserCreate struct {
	Name     string
	Nickname string
	Birthday time.Time
	Phone    int16
	Email    string
	Password string
	State    bool
}
