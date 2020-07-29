package users

import (
	"database/sql"
	"github.com/ernesto2108/APIRest_Business/internal/logs"
	entities "github.com/ernesto2108/APIRest_Business/internal/pkg/users/domain"
)

func (us UserServices) create(u *entities.UserCreate) (*entities.User, error) {

	//	use services postgres
	tx, err := us.SqlClient.Begin()

	if err != nil {
		logs.Log().Error("Cannot create transaction")
		return nil, err
	}

	//	data definition null

	intNull := sql.NullInt64{}

	if u.Phone == 0 {
		intNull.Valid = false
	} else {
		intNull.Valid = true
		intNull.Int64 = int64(u.Phone)
	}

	if u.IdentityDocument == 0 {
		intNull.Valid = false
	} else {
		intNull.Valid = true
		intNull.Int64 = int64(u.IdentityDocument)
	}

	stringNull := sql.NullString{}

	if u.Name == 0 {
		stringNull.Valid = false
	} else {
		stringNull.Valid = true
		stringNull.String = u.Name
	}

	if u.Nickname == 0 {
		stringNull.Valid = false
	} else {
		stringNull.Valid = true
		stringNull.String = u.Nickname
	}

	if u.Email == 0 {
		stringNull.Valid = false
	} else {
		stringNull.Valid = true
		stringNull.String = u.Email
	}

	if u.Password == 0 {
		stringNull.Valid = false
	} else {
		stringNull.Valid = true
		stringNull.String = u.Password
	}

	bollNull := sql.NullBool{}

	if u.State == 0 {
		bollNull.Valid = false
	} else {
		bollNull.Valid = true
		bollNull.Bool = u.State
	}

	//	insert data to database
	query, err := tx.Exec(`INSERT INTO users (first_name, last_name, birthday, phone, email, password, state),
			VALUES ($1, $2, $3, $4, $5, $6, $7)`, stringNull, stringNull, intNull, intNull, stringNull, stringNull, bollNull)

	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	//	rows affected
	id, err := query.RowsAffected()

	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}

	_ = tx.Commit()

	//	return user entities
	return &entities.User{
		Id:               id,
		FirstName:        u.FirstName,
		LastName:         u.LastName,
		IdentityDocument: u.IdentityDocument,
		Phone:            u.Phone,
		Email:            u.Email,
		Password:         u.Password,
		State:            u.State,
	}
}
