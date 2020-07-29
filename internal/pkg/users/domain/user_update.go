package users

import (
	"github.com/google/uuid"
)

type UserUpdate struct {
	Id               uuid.UUID `json:"id"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	IdentityDocument int16     `json:"identity_document"`
	Phone            int16     `json:"phone"`
	Email            string    `json:"email"`
	Password         string    `json:"password"`
	State            bool      `json:"state"`
}
