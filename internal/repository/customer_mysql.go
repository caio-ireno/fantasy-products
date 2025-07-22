package repository

import (
	"app/internal/domain"
	"database/sql"
	"fmt"
)

// NewCustomersMySQL creates new mysql repository for customer entity.
func NewCustomersMySQL(db *sql.DB) *CustomersMySQL {
	return &CustomersMySQL{db}
}

// CustomersMySQL is the MySQL repository implementation for customer entity.
type CustomersMySQL struct {
	// db is the database connection.
	db *sql.DB
}

// FindAll returns all customers from the database.
func (r *CustomersMySQL) FindAll() (c []domain.Customer, err error) {
	// execute the query
	rows, err := r.db.Query("SELECT `id`, `first_name`, `last_name`, `condition` FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var cs domain.Customer
		// scan the row into the customer
		err := rows.Scan(&cs.Id, &cs.FirstName, &cs.LastName, &cs.Condition)
		if err != nil {
			return nil, err
		}
		// append the customer to the slice
		c = append(c, cs)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

// Save saves the customer into the database.
func (r *CustomersMySQL) Save(c *domain.Customer) (err error) {
	res, err := r.db.Exec(
		"INSERT INTO customers (`first_name`, `last_name`, `condition`) VALUES (?, ?, ?)",
		(*c).FirstName, (*c).LastName, (*c).Condition,
	)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	(*c).Id = int(id)
	fmt.Print((*c).Id)
	return
}

// Save saves the customer into the database.
func (r *CustomersMySQL) SaveJson(c []*domain.Customer) (total int, err error) {

	for _, customer := range c {
		total++
		_, err = r.db.Exec("INSERT INTO customers (`first_name`, `last_name`, `condition`) VALUES (?, ?, ?)",
			(*customer).CustomerAttributes.FirstName, (*customer).CustomerAttributes.LastName, (*customer).CustomerAttributes.Condition,
		)

		if err != nil {
			return
		}

	}
	return
}

// req 1 Valores totais arredondados para 2 casas decimais por condition do customer
func (r *CustomersMySQL) GetTotalByCondition() (d []domain.CustomerGetTotal, err error) {

	rows, err := r.db.Query(`SELECT 
    CASE c.condition
        WHEN 1 THEN 'Activo ( 1 )'
        ELSE 'Inactivo ( 0 )'
    END AS Condition1,
    ROUND(SUM(i.total), 2) AS Total
FROM invoices i
JOIN customers c ON c.id = i.customer_id
GROUP BY c.condition;`)

	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var totalCustomer domain.CustomerGetTotal
		err = rows.Scan(&totalCustomer.Condition, &totalCustomer.Total)
		if err != nil {
			return
		}
		d = append(d, totalCustomer)
	}
	err = rows.Err()
	return

}

// req 3 Top 5 dos customers ativos quem gastou mais dinheiro
func (r *CustomersMySQL) GetMostActive() (ma []domain.CustomerGetMostActive, err error) {

	rows, err := r.db.Query(`select c.first_name, c.last_name, round(sum( i.total),2) as total
from invoices i
join customers c on i.customer_id=c.id
group by customer_id
order by total desc
limit 5;`)

	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var totalCustomer domain.CustomerGetMostActive
		err = rows.Scan(&totalCustomer.FirstName, &totalCustomer.LastName, &totalCustomer.Amount)
		if err != nil {
			return
		}
		ma = append(ma, totalCustomer)
	}
	err = rows.Err()
	return

}
