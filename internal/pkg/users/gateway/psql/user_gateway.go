package users

import (
	entities "github.com/ernesto2108/APIRest_Business/internal/pkg/users/domain"
	"github.com/ernesto2108/APIRest_Business/internal/services/psql"
)

type UserGateway interface {
	Create(u *entities.UserCreate) (*entities.User, error)
	FindId(id int64) (*entities.User, error)
	Find() []*entities.User
}

type UserCreateGateway struct {
	UserServicesGateway
}

func NewUserCreateGateway(s *services.SqlClient) UserCreateGateway {
	return &UserCreateGateway{
		&UserServices{s},
	}
}
