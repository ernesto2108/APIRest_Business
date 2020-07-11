package menu

type MenuCreateGateway interface {
	create()
	get()
	getId()
	update()
	delete()
}

type MenuCreateGtw struct {
}

func NewMenuCreateGateway()  {
}