package business

type BusinessGateway interface {
	create()
	get()
	getId()
	update()
	delete()
}

type BusinessCreateGtw struct {
}

func NewBusinessCreateGateway() {
}