package repository_test

import (
	"app/internal/domain"
	"app/internal/repository"
	"app/tests"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSalesRepository(t *testing.T) {
	tests.RegisterDatabase()
	tests.InitDatabase()

	db := tests.GetDB()
	repo := repository.NewSalesMySQL(db)

	t.Run("should get the total by invoices and customer id", func(t *testing.T) {
		body := []domain.SaleTopFiveProducts([]domain.SaleTopFiveProducts{domain.SaleTopFiveProducts{Description: "Vinegar - Raspberry", Total: 660}, domain.SaleTopFiveProducts{Description: "Flour - Corn, Fine", Total: 521}, domain.SaleTopFiveProducts{Description: "Cookie - Oatmeal", Total: 467}, domain.SaleTopFiveProducts{Description: "Pepper - Red Chili", Total: 439}, domain.SaleTopFiveProducts{Description: "Chocolate - Milk Coating", Total: 436}})
		data, err := repo.GetTopFiveProducts()

		assert.Equal(t, body, data)
		assert.Nil(t, err)
	})

}
