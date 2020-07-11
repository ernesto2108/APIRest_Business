package products

import (
	"github.com/google/uuid"

	business "github.com/ernesto2108/APIRest_Business/internal/pkg/business/domain"
	category "github.com/ernesto2108/APIRest_Business/internal/pkg/categories/domain"
)

type Product struct {
	ID          uuid.UUID
	Image       string
	Title       string
	Description string
	Price       string
	Quality     string
	Category   []category.Category
	Business   []business.Business
}
