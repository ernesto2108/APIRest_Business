package psql

type UserCreateGateway interface {
	create()
	get()
	getId()
	update()
	delete()
}

type UserCreateGtw struct {
}

func NewUserCreateGateway()  {
}
