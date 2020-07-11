package category

import "github.com/google/uuid"

type Category struct {
	Id    uuid.UUID
	Image string
	Name  string
}
