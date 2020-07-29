package users

import (
	"github.com/ernesto2108/APIRest_Business/internal/logs"
	entities "github.com/ernesto2108/APIRest_Business/internal/pkg/users/domain"
)

func (us UserServices) delete(id int64) *entities.User {

	//	use services postgres
	tx, err := us.SqlClient.Begin()

	if err != nil {
		logs.Log().Error("cannot execute statement")
		return nil
	}

	var user entities.User

	//	delete user data
	err = tx.QueryRow(`SELECT id, first_name, last_name, phone, identity_document, email, password, 
			state FROM users WHERE id = $1`).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Phone,
		&user.IdentityDocument, &user.Email, &user.Password, &user.State)

	if err != nil {
		logs.Log().Error("Unable to execute query")
		_ = tx.Rollback()
		return nil
	}

	_, err = tx.Exec("DELETE FORM users WHERE id = $1", id)

	if err != nil {
		logs.Log().Error("Unable to execute query")
		_ = tx.Rollback()
		return nil
	}

	_ = tx.Commit()

	return &entities.User{
		Id:               id,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		Phone:            user.Phone,
		IdentityDocument: user.IdentityDocument,
		Email:            user.Email,
		Password:         user.Password,
		State:            user.State,
	}
}
