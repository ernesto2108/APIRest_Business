package users

import (
	"github.com/ernesto2108/APIRest_Business/internal/logs"
	entities "github.com/ernesto2108/APIRest_Business/internal/pkg/users/domain"
	channel "github.com/ernesto2108/APIRest_Business/internal/util/channels"
)

func (us UserServices) find() ([]*entities.User, error)  {

	//	use services postgres
	tx, err := us.SqlClient.Begin()

	if err != nil {
		logs.Log().Error("Cannot create transaction")
		return nil, err
	}

	var users []*entities.User

	done := make(chan bool)

	go func(ch chan<- bool) {

		//	find all data of users
		query, err := tx.Query(`SELECT id, first_name, last_name, phone, dni, email, password,
 		state FROM users`)

		if err != nil {
			logs.Log().Error("Cannot search statement")
			_ = tx.Rollback()
			return nil, err
		}

		defer query.Close()
		
		for query.Next() {
			var user entities.User
			err := query.Scan(&user.FirstName, &user.LastName, &user.Phone, &user.IdentityDocument,
				&user.Email, &user.Password, &user.State)
			if err != nil {
				logs.Log().Error("Cannot read current row")
				return nil
			}
			users = append(users, &user)
		}

		ch <- true

	}(done)

	if channel.OK(done) {
		return users, nil
	}

	return nil, err
}
