package service

import "app/internal/domain"

// NewSalesDefault creates new default service for sale entity.
func NewSalesDefault(rp domain.RepositorySale) *SalesDefault {
	return &SalesDefault{rp}
}

// SalesDefault is the default service implementation for sale entity.
type SalesDefault struct {
	// rp is the repository for sale entity.
	rp domain.RepositorySale
}

// FindAll returns all sales.
func (sv *SalesDefault) FindAll() (s []domain.Sale, err error) {
	s, err = sv.rp.FindAll()
	return
}

// Save saves the sale.
func (sv *SalesDefault) Save(s *domain.Sale) (err error) {
	err = sv.rp.Save(s)
	return
}

// Save saves the customer.
func (s *SalesDefault) SaveJson(sa []*domain.Sale) (total int, err error) {
	total, err = s.rp.SaveJson(sa)
	return
}
