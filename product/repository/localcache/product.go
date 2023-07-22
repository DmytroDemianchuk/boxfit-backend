package localcache

import (
	"context"
	"sync"

	"github.com/dmytrodemianchuk/boxfit-backend/models"
	"github.com/dmytrodemianchuk/boxfit-backend/product"
)

type ProductLocalStorage struct {
	products map[string]*models.Product
	mutex    *sync.Mutex
}

func NewProductLocalStorage() *ProductLocalStorage {
	return &ProductLocalStorage{
		products: make(map[string]*models.Product),
		mutex:    new(sync.Mutex),
	}
}

func (s *ProductLocalStorage) CreateProduct(ctx context.Context, user *models.User, bm *models.Product) error {
	bm.UserID = user.ID

	s.mutex.Lock()
	s.products[bm.ID] = bm
	s.mutex.Unlock()

	return nil
}

func (s *ProductLocalStorage) GetProducts(ctx context.Context, user *models.User) ([]*models.Product, error) {
	products := make([]*models.Product, 0)

	s.mutex.Lock()
	for _, bm := range s.products {
		if bm.UserID == user.ID {
			products = append(products, bm)
		}
	}

	s.mutex.Unlock()

	return products, nil
}

func (s *ProductLocalStorage) DeleteProduct(ctx context.Context, user *models.User, id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	bm, ex := s.products[id]
	if ex && bm.UserID == user.ID {
		delete(s.products, id)
		return nil
	}

	return product.ErrProductNotFound
}
