package users

import (
	"github.com/ernesto2108/APIRest_Business/internal/logs"
	entities "github.com/ernesto2108/APIRest_Business/internal/pkg/users/domain"
)

func (us UserServices) update(u *entities.UserUpdate) *entities.User {

	//
	tx, err := us.SqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot execute statement")
		return nil
	}

	//	update user data
	_, err = tx.Exec(`UPDATE users SET name = $1, nickname = $2, phone = $3, identity_document = $4, 
			email = $5, password = $6, state = $ WHERE id = $7`, u.FirstName, u.LastName, u.Phone,
		u.IdentityDocument, u.Email, u.Password, u.State)

	if err != nil {
		logs.Log().Error("Unable to execute query")
		_ = tx.Rollback()
		return nil
	}

	_ = tx.Commit()

	return &entities.User{
		Id:               u.Id,
		FirstName:        u.FirstName,
		LastName:         u.LastName,
		Phone:            u.Phone,
		IdentityDocument: u.IdentityDocument,
		Email:            u.Email,
		Password:         u.Password,
		State:            u.State,
	}
}
