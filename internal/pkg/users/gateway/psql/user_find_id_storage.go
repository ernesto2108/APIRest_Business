package users

import (
	"github.com/ernesto2108/APIRest_Business/internal/logs"
	entities "github.com/ernesto2108/APIRest_Business/internal/pkg/users/domain"
	channel "github.com/ernesto2108/APIRest_Business/internal/util/channels"
)

func (us UserServices) findId(id int64) (*entities.User, error) {

	//	use services postgres
	tx, err := us.SqlClient.Begin()

	if err != nil {
		logs.Log().Error("Cannot create transaction")
		return nil, err
	}

	var u entities.User

	done := make(chan bool)

	go func(ch chan<- bool) {

		//	find by id the users
		err = tx.QueryRow(`SELECT first_name, last_name, phone, identity_document, email, password,
 		state FROM users WHERE id = $1`, id).Scan(&u.FirstName, &u.LastName, &u.Phone, &u.IdentityDocument,
			&u.Email, &u.Password, &u.State)

		if err != nil {
			ch <- false
			logs.Log().Error("cannot execute statement")
			return nil, err
		}

		ch <- true

	}(done)

	if channel.OK(done) {
		return u, nil
	}

	return &u, nil
}
