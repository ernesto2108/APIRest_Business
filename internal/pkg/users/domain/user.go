package users

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id       uuid.UUID
	Name     string
	Nickname string
	Birthday time.Time
	Phone    int16
	Email    string
	Password string
	State    bool
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
}
