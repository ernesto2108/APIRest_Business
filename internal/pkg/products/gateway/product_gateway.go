package products

type ProductCreateGateway interface {
	create()
	get()
	getId()
	update()
	delete()
}

type ProductCreateGtw struct {
}

func NewProductCreateGateway()  {
}