package handler

import (
	"app/internal/domain"
	"app/utils"
	"fmt"
	"net/http"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
)

// NewInvoicesDefault returns a new InvoicesDefault
func NewInvoicesDefault(sv domain.ServiceInvoice) *InvoicesDefault {
	return &InvoicesDefault{sv: sv}
}

// InvoicesDefault is a struct that returns the invoice handlers
type InvoicesDefault struct {
	// sv is the invoice's service
	sv domain.ServiceInvoice
}

// InvoiceJSON is a struct that represents a invoice in JSON format
type InvoiceJSON struct {
	Id         int     `json:"id"`
	Datetime   string  `json:"datetime"`
	Total      float64 `json:"total"`
	CustomerId int     `json:"customer_id"`
}

// GetAll returns all invoices
func (h *InvoicesDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// ...

		// process
		i, err := h.sv.FindAll()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error getting invoices")
			return
		}

		// response
		// - serialize
		ivJSON := make([]InvoiceJSON, len(i))
		for ix, v := range i {
			ivJSON[ix] = InvoiceJSON{
				Id:         v.Id,
				Datetime:   v.Datetime,
				Total:      v.Total,
				CustomerId: v.CustomerId,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "invoices found",
			"data":    ivJSON,
		})
	}
}

// RequestBodyInvoice is a struct that represents the request body for a invoice
type RequestBodyInvoice struct {
	Datetime   string  `json:"datetime"`
	Total      float64 `json:"total"`
	CustomerId int     `json:"customer_id"`
}

// Create creates a new invoice
func (h *InvoicesDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// request
		// - body
		var reqBody RequestBodyInvoice
		err := request.JSON(r, &reqBody)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "error parsing request body")
			return
		}

		// process
		// - deserialize
		i := domain.Invoice{
			InvoiceAttributes: domain.InvoiceAttributes{
				Datetime:   reqBody.Datetime,
				Total:      reqBody.Total,
				CustomerId: reqBody.CustomerId,
			},
		}
		// - save
		err = h.sv.Save(&i)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error saving invoice")
			return
		}

		// response
		// - serialize
		iv := InvoiceJSON{
			Id:         i.Id,
			Datetime:   i.Datetime,
			Total:      i.Total,
			CustomerId: i.CustomerId,
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "invoice created",
			"data":    iv,
		})
	}
}

// Create creates a new customer
func (h *InvoicesDefault) CreateWithJson() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		invoices, err := utils.ReadJson[InvoiceJSON]("invoices")
		fmt.Println(invoices)
		if err != nil {
			response.Error(w, http.StatusBadRequest, "error get json data to file")

		}

		var invoicesPtrs []*domain.Invoice
		for i := range invoices {
			invoices := &domain.Invoice{
				InvoiceAttributes: domain.InvoiceAttributes{
					Datetime:   invoices[i].Datetime,
					CustomerId: invoices[i].CustomerId,
					Total:      invoices[i].Total,
				},
			}
			invoicesPtrs = append(invoicesPtrs, invoices)
		}

		total, err := h.sv.SaveJson(invoicesPtrs)

		if err != nil {
			response.Error(w, http.StatusInternalServerError, "error saving customer")
			return
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "invoices created",
			"data":    total,
		})
	}
}
