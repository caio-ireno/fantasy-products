package repository

import (
	"app/internal/domain"
	"database/sql"
)

// NewInvoicesMySQL creates new mysql repository for invoice entity.
func NewInvoicesMySQL(db *sql.DB) *InvoicesMySQL {
	return &InvoicesMySQL{db}
}

// InvoicesMySQL is the MySQL repository implementation for invoice entity.
type InvoicesMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all invoices from the database.
func (r *InvoicesMySQL) FindAll() (i []domain.Invoice, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `datetime`, `total`, `customer_id` FROM invoices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var iv domain.Invoice
		// scan the row into the invoice
		err := rows.Scan(&iv.Id, &iv.Datetime, &iv.Total, &iv.CustomerId)
		if err != nil {
			return nil, err
		}
		// append the invoice to the slice
		i = append(i, iv)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the invoice into the database.
func (r *InvoicesMySQL) Save(i *domain.Invoice) (err error) {
	// execute the query
	res, err := r.db.Exec(
		"INSERT INTO invoices (`datetime`, `total`, `customer_id`) VALUES (?, ?, ?)",
		(*i).Datetime, (*i).Total, (*i).CustomerId,
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
	(*i).Id = int(id)

	return
}

// Save saves the customer into the database.
func (r *InvoicesMySQL) SaveJson(c []*domain.Invoice) (total int, err error) {

	for _, invoice := range c {
		total++
		_, err = r.db.Exec("INSERT INTO invoices (`datetime`, `total`, `customer_id`) VALUES (?, ?, ?)",
			(*invoice).Datetime, (*invoice).Total, (*invoice).CustomerId)

		if err != nil {
			return
		}

	}
	return
}

func (r *InvoicesMySQL) GetTotalByInvoicesIdAndCustomerId() (t []domain.InvoiceTotalToUpdate, err error) {

	rows, err := r.db.Query(`select SUM(s.quantity) as total, s.invoice_id
		from 
		sales s 
		join invoices i on s.invoice_id=i.id
		join customers c on c.id =i.customer_id
		join products p on s.product_id= p.id
		GROUP BY s.invoice_id`,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var totalToUpdate domain.InvoiceTotalToUpdate
		err = rows.Scan(&totalToUpdate.Total, &totalToUpdate.InvoiceId)
		if err != nil {
			return
		}
		t = append(t, totalToUpdate)
	}
	err = rows.Err()
	return
}

func (r *InvoicesMySQL) UpdateTotal(data []domain.InvoiceTotalToUpdate) (err error) {

	for _, d := range data {
		query := "update invoices set total = ROUND(?,2) where id=?"
		_, err = r.db.Exec(query, d.Total, d.InvoiceId)

		if err != nil {
			return
		}
	}
	return
}
