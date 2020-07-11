package users

import (
	"github.com/google/uuid"
	"time"
)

type UserUpdate struct {
	Id       uuid.UUID
	Name     string
	Nickname string
	Birthday time.Time
	Phone    int16
	Email    string
	Password string
	State    bool
}
