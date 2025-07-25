package service

import "app/internal/domain"

// NewInvoicesDefault creates new default service for invoice entity.
func NewInvoicesDefault(rp domain.RepositoryInvoice) *InvoicesDefault {
	return &InvoicesDefault{rp}
}

// InvoicesDefault is the default service implementation for invoice entity.
type InvoicesDefault struct {
	// rp is the repository for invoice entity.
	rp domain.RepositoryInvoice
}

// FindAll returns all invoices.
func (s *InvoicesDefault) FindAll() (i []domain.Invoice, err error) {
	i, err = s.rp.FindAll()
	return
}

// Save saves the invoice.
func (s *InvoicesDefault) Save(i *domain.Invoice) (err error) {
	err = s.rp.Save(i)
	return
}

// Save saves the customer.
func (s *InvoicesDefault) SaveJson(c []*domain.Invoice) (total int, err error) {
	total, err = s.rp.SaveJson(c)
	return
}

func (s *InvoicesDefault) UpdateTotal() (err error) {
	data, err := s.rp.GetTotalByInvoicesIdAndCustomerId()
	if err != nil {
		return
	}

	err = s.rp.UpdateTotal(data)

	if err != nil {
		return
	}

	return
}
