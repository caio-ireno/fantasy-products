package handler

import (
	"app/internal/domain"
	"app/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
)

// NewCustomersDefault returns a new CustomersDefault
func NewCustomersDefault(sv domain.ServiceCustomer) *CustomersDefault {
	return &CustomersDefault{sv: sv}
}

// CustomersDefault is a struct that returns the customer handlers
type CustomersDefault struct {
	// sv is the customer's service
	sv domain.ServiceCustomer
}

// CustomerJSON is a struct that represents a customer in JSON format
type CustomerJSON struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Condition int    `json:"condition"`
}

// RequestBodyCustomer is a struct that represents the request body for a customer
type RequestBodyCustomer struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Condition int    `json:"condition"`
}

// GetAll returns all customers
func (h *CustomersDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		c, err := h.sv.FindAll()
		if err != nil {
			log.Println(err)
			response.Error(w, http.StatusInternalServerError, "error getting customers")
			return
		}

		// response
		// - serialize
		csJSON := make([]CustomerJSON, len(c))
		for ix, v := range c {
			csJSON[ix] = CustomerJSON{
				Id:        v.Id,
				FirstName: v.FirstName,
				LastName:  v.LastName,
				Condition: v.Condition,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "customers found",
			"data":    csJSON,
		})
	}
}

// Create creates a new customer
func (h *CustomersDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var reqBody RequestBodyCustomer
		err := request.JSON(r, &reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "error deserializing request body")
			return
		}

		c := domain.Customer{
			CustomerAttributes: domain.CustomerAttributes{
				FirstName: reqBody.FirstName,
				LastName:  reqBody.LastName,
				Condition: reqBody.Condition,
			},
		}
		// - save
		err = h.sv.Save(&c)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error saving customer")
			return
		}

		// response
		// - serialize
		cs := CustomerJSON{
			Id:        c.Id,
			FirstName: c.FirstName,
			LastName:  c.LastName,
			Condition: c.Condition,
		}
		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "customer created",
			"data":    cs,
		})
	}
}

// Create creates a new customer
func (h *CustomersDefault) CreateWithJson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		customers, err := utils.ReadJson[CustomerJSON]("customers")
		fmt.Println(customers)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "error get json data to file")

		}

		var customerPtrs []*domain.Customer
		for i := range customers {
			customer := &domain.Customer{
				CustomerAttributes: domain.CustomerAttributes{
					FirstName: customers[i].FirstName,
					LastName:  customers[i].LastName,
					Condition: customers[i].Condition,
				},
			}
			customerPtrs = append(customerPtrs, customer)
		}

		total, err := h.sv.SaveJson(customerPtrs)

		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error saving customer")
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "customer created",
			"data":    total,
		})
	}
}

// Create creates a new customer
func (h *CustomersDefault) GetTotalByCondition() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		total, err := h.sv.GetTotalByCondition()

		if err != nil {
			fmt.Println(err)
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Total  by condition",
			"data":    total,
		})
	}
}

func (h *CustomersDefault) GetMostActive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ma, err := h.sv.GetMostActive()

		if err != nil {
			response.Error(w, http.StatusInternalServerError, err.Error())
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Most active customer spent more money",
			"data":    ma,
		})
	}
}
