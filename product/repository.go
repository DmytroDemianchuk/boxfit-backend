package product

import (
	"context"

	"github.com/dmytrodemianchuk/boxfit-backend/models"
)

type Repository interface {
	CreateProduct(ctx context.Context, user *models.User, bm *models.Product) error
	GetProducts(ctx context.Context, user *models.User) ([]*models.Product, error)
	DeleteProduct(ctx context.Context, user *models.User, id string) error
}
