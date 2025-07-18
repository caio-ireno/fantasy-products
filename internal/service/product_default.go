package service

import "app/internal/domain"

// NewProductsDefault creates new default service for product entity.
func NewProductsDefault(rp domain.RepositoryProduct) *ProductsDefault {
	return &ProductsDefault{rp}
}

// ProductsDefault is the default service implementation for product entity.
type ProductsDefault struct {
	// rp is the repository for product entity.
	rp domain.RepositoryProduct
}

// FindAll returns all products.
func (s *ProductsDefault) FindAll() (p []domain.Product, err error) {
	p, err = s.rp.FindAll()
	return
}

// Save saves the product.
func (s *ProductsDefault) Save(p *domain.Product) (err error) {
	err = s.rp.Save(p)
	return
}

// Save saves the customer.
func (s *ProductsDefault) SaveJson(c []*domain.Product) (total int, err error) {
	total, err = s.rp.SaveJson(c)
	return
}
