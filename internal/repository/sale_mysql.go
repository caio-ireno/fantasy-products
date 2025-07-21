package repository

import (
	"app/internal/domain"
	"database/sql"
)

// NewSalesMySQL creates new mysql repository for sale entity.
func NewSalesMySQL(db *sql.DB) *SalesMySQL {
	return &SalesMySQL{db}
}

// SalesMySQL is the MySQL repository implementation for sale entity.
type SalesMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all sales from the database.
func (r *SalesMySQL) FindAll() (s []domain.Sale, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `quantity`, `product_id`, `invoice_id` FROM sales")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var sa domain.Sale
		// scan the row into the sale
		err := rows.Scan(&sa.Id, &sa.Quantity, &sa.ProductId, &sa.InvoiceId)
		if err != nil {
			return nil, err
		}
		// append the sale to the slice
		s = append(s, sa)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the sale into the database.
func (r *SalesMySQL) Save(s *domain.Sale) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO sales (`quantity`, `product_id`, `invoice_id`) VALUES (?, ?, ?)",
		(*s).Quantity, (*s).ProductId, (*s).InvoiceId,
	)
	if err != nil {
		return err
	}

	// get the last inserted id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set the id
	(*s).Id = int(id)

	return
}

// Save saves the customer into the database.
func (r *SalesMySQL) SaveJson(c []*domain.Sale) (total int, err error) {

	for _, s := range c {
		total++
		_, err = r.db.Exec(
			"INSERT INTO sales (`quantity`, `product_id`, `invoice_id`) VALUES (?, ?, ?)",
			(*s).Quantity, (*s).ProductId, (*s).InvoiceId,
		)

		if err != nil {
			return
		}

	}
	return
}

func (r *SalesMySQL) GetTopFiveProducts() (s []domain.SaleTopFiveProducts, err error) {

	rows, err := r.db.Query(`select p.description, sum(s.quantity) as total  from 
							sales s
							join products p on p.id=s.product_id
							group by product_id
							order by total desc
							limit 5;`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var sa domain.SaleTopFiveProducts
		// scan the row into the sale
		err := rows.Scan(&sa.Description, &sa.Total)
		if err != nil {
			return nil, err
		}
		// append the sale to the slice
		s = append(s, sa)
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return
}
