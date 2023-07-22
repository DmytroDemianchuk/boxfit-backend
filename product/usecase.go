package product

import (
	"context"

	"github.com/dmytrodemianchuk/boxfit-backend/models"
)

type UseCase interface {
	CreateProduct(ctx context.Context, user *models.User, name, category string) error
	GetProducts(ctx context.Context, user *models.User) ([]*models.Product, error)
	DeleteProduct(ctx context.Context, user *models.User, id string) error
}
