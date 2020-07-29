package users

import (
	entities "github.com/ernesto2108/APIRest_Business/internal/pkg/users/domain"
	"github.com/ernesto2108/APIRest_Business/internal/services/psql"
)

type UserServicesGateway interface {
	create(u *entities.UserCreate) (*entities.User, error)
	findId(id int64) (*entities.User, error)
	find() ([]*entities.User, error)
	update(u *entities.UserUpdate) *entities.User
	delete(id int64) *entities.User
}

type UserServices struct {
	*services.SqlClient
}

