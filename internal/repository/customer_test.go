package repository_test

import (
	"app/internal/domain"
	"app/internal/repository"
	"app/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomerRepository(t *testing.T) {
	tests.RegisterDatabase()
	tests.InitDatabase()

	db := tests.GetDB()
	repo := repository.NewCustomersMySQL(db)

	t.Run("should save customer", func(t *testing.T) {
		// given

		customer := domain.Customer{
			CustomerAttributes: domain.CustomerAttributes{
				FirstName: "Lannie",
				LastName:  "Test",
				Condition: 1,
			},
			Id: 1,
		}

		// when
		err := repo.Save(&customer)

		// then
		assert.Nil(t, err)
		assert.NotEqual(t, customer.Id, 0)

	})

	t.Run("should save customer with json", func(t *testing.T) {
		customerJson := []*domain.Customer{
			{
				CustomerAttributes: domain.CustomerAttributes{
					FirstName: "Lannie",
					LastName:  "Test",
					Condition: 1,
				},
				Id: 1,
			},
			{
				CustomerAttributes: domain.CustomerAttributes{
					FirstName: "Bonnie",
					LastName:  "Test",
					Condition: 0,
				},
				Id: 2,
			},
		}

		total, err := repo.SaveJson(customerJson)

		assert.Nil(t, err)
		assert.NotEqual(t, total, 0)
		assert.Equal(t, total, 2)

	})

	t.Run("Should get total by condition", func(t *testing.T) {
		body := []domain.CustomerGetTotal{{Condition: "Inactivo ( 0 )", Total: 605929.11}, {Condition: "Activo ( 1 )", Total: 716792.33}}

		data, err := repo.GetTotalByCondition()

		assert.Equal(t, body, data)
		assert.Nil(t, err)

	})

	t.Run("Should find all. Customer", func(t *testing.T) {
		_, err := db.Exec("DELETE FROM customers")
		assert.Nil(t, err)

		customers := []domain.Customer{
			{CustomerAttributes: domain.CustomerAttributes{FirstName: "A", LastName: "B", Condition: 1}},
			{CustomerAttributes: domain.CustomerAttributes{FirstName: "C", LastName: "D", Condition: 0}},
		}
		for i := range customers {
			err := repo.Save(&customers[i])
			assert.Nil(t, err)
		}

		data, err := repo.FindAll()

		assert.Equal(t, len(customers), len(data))
		assert.Nil(t, err)
	})

	t.Run("Should get most active spent money", func(t *testing.T) {
		db.Exec("DELETE FROM customers")
		db.Exec("DELETE FROM products")
		db.Exec("DELETE FROM invoices")
		db.Exec("DELETE FROM sales")

		_, err := db.Exec("INSERT INTO products (id, description, price) VALUES (1, 'Produto1', 100.0)")
		assert.Nil(t, err)
		_, err = db.Exec("INSERT INTO customers (id, first_name, last_name, `condition`) VALUES (1, 'Lannie', 'Tortis', 1)")
		assert.Nil(t, err)
		_, err = db.Exec("INSERT INTO invoices (id, customer_id, total) VALUES (1, 1, 100.0)")
		assert.Nil(t, err)
		_, err = db.Exec("INSERT INTO sales (id, invoice_id, product_id, quantity) VALUES (1, 1, 1, 1)")
		assert.Nil(t, err)

		data, err := repo.GetMostActive()
		assert.Nil(t, err)
		assert.NotEmpty(t, data)
		if len(data) > 0 {
			assert.Equal(t, "Lannie", data[0].FirstName)
		}
	})
}
