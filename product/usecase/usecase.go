package usecase

import (
	"context"
	"time"

	"github.com/dmytrodemianchuk/boxfit-backend/models"
	"github.com/dmytrodemianchuk/boxfit-backend/product"
)

type ProductUseCase struct {
	productRepo product.Repository
}

func NewProducUseCase(productRepo product.Repository) *ProductUseCase {
	return &ProductUseCase{
		productRepo: productRepo,
	}
}

func (p ProductUseCase) CreateProduct(ctx context.Context, user *models.User, name, category, subcategory, mark, variant, color string, image []byte, number uint16, price uint32, created_at, updated_at time.Time) error {
	pr := &models.Product{
		Name:        name,
		Category:    category,
		Subcategory: subcategory,
		Mark:        mark,
		Variant:     variant,
		Color:       color,
		Number:      number,
		Price:       price,
		Image:       image,
		CreatedAt:   created_at,
		UpdatedAt:   updated_at,
	}

	return p.productRepo.CreateProduct(ctx, user, pr)
}

func (p ProductUseCase) GetProduct(ctx context.Context, user *models.User) ([]*models.Product, error) {
	return p.productRepo.GetProducts(ctx, user)
}

func (p ProductUseCase) DeleteProduct(ctx context.Context, user *models.User, id string) error {
	return p.productRepo.DeleteProduct(ctx, user, id)
}
