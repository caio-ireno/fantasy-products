package domain

// RepositoryCustomer is the interface that wraps the basic methods that a customer repository should implement.
type RepositoryCustomer interface {
	// FindAll returns all customers saved in the database.
	FindAll() (c []Customer, err error)
	// Save saves a customer into the database.
	Save(c *Customer) (err error)
	SaveJson(c []*Customer) (total int, err error)
	GetTotalByCondition() (d []CustomerGetTotal, err error)

	GetMostActive() (ma []CustomerGetMostActive, err error)
}

// ServiceCustomer is the interface that wraps the basic methods that a customer service should implement.
type ServiceCustomer interface {
	// FindAll returns all customers
	FindAll() (c []Customer, err error)
	// Save saves a customer
	Save(c *Customer) (err error)
	SaveJson(c []*Customer) (total int, err error)
	GetTotalByCondition() (d []CustomerGetTotal, err error)
	GetMostActive() (ma []CustomerGetMostActive, err error)
}

// CustomerAttributes is the struct that represents the attributes of a customer.
type CustomerAttributes struct {
	// FirstName is the first name of the customer.
	FirstName string
	// LastName is the last name of the customer.
	LastName string
	// Condition is the condition of the customer.
	Condition int
}

// Customer is the struct that represents a customer.
type Customer struct {
	// Id is the unique identifier of the customer.
	Id int
	// CustomerAttributes is the attributes of the customer.
	CustomerAttributes
}

type CustomerGetTotal struct {
	Condition string
	Total     float64
}

type CustomerGetMostActive struct {
	FirstName string
	LastName  string
	Amount    float64
}
