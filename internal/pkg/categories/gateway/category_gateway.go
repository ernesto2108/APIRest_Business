package categories

type CategoryCreateGateway interface {
	create()
	get()
	getId()
	update()
	delete()
}

type CategoryCreateGtw struct {
}

func NewCategoryCreateGateway()  {
}
