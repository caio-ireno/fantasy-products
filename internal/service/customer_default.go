package service

import "app/internal/domain"

func NewCustomersDefault(rp domain.RepositoryCustomer) *CustomersDefault {
	return &CustomersDefault{rp}
}

type CustomersDefault struct {
	// rp is the repository for customer entity.
	rp domain.RepositoryCustomer
}

func (s *CustomersDefault) FindAll() (c []domain.Customer, err error) {
	c, err = s.rp.FindAll()
	return
}

func (s *CustomersDefault) Save(c *domain.Customer) (err error) {
	err = s.rp.Save(c)
	return
}

func (s *CustomersDefault) SaveJson(c []*domain.Customer) (total int, err error) {
	total, err = s.rp.SaveJson(c)
	return
}

// Save saves the customer.
func (s *CustomersDefault) GetTotalByCondition() (d []domain.CustomerGetTotal, err error) {
	d, err = s.rp.GetTotalByCondition()
	return
}
