package handler_test

import (
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"app/tests"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func setupCustomerHandler() *handler.CustomersDefault {
	tests.RegisterDatabase()
	tests.InitDatabase()
	db := tests.GetDB()
	repo := repository.NewCustomersMySQL(db)
	svc := service.NewCustomersDefault(repo)
	return handler.NewCustomersDefault(svc)
}

func TestGetTotalByCondition_Integration(t *testing.T) {
	h := setupCustomerHandler()
	req := httptest.NewRequest(http.MethodGet, "/customers/totalByCondition", nil)
	w := httptest.NewRecorder()
	h.GetTotalByCondition()(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Decodifica e valida o body
	var body map[string]any
	err := json.NewDecoder(resp.Body).Decode(&body)
	assert.Equal(t, nil, err)
	assert.Equal(t, "Total  by condition", body["message"])
	data, ok := body["data"].([]any)
	fmt.Println(data)
	if ok {
		assert.Equal(t, true, len(data) > 0)
	} else {
		t.Errorf("expected data to be a map, got %T", body["data"])
	}
}

func TestGetMostActive_Integration(t *testing.T) {
	h := setupCustomerHandler()
	req := httptest.NewRequest(http.MethodGet, "/customers/mostActive", nil)
	w := httptest.NewRecorder()
	h.GetMostActive()(w, req)
	resp := w.Result()
	fmt.Println(resp)
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var body map[string]any
	err := json.NewDecoder(resp.Body).Decode(&body)
	assert.Equal(t, nil, err)
	assert.Equal(t, "Most active customer spent more money", body["message"])
	fmt.Println(body)
	if data, ok := body["data"].([]any); ok {
		if len(data) > 0 {
			if first, ok := data[0].(map[string]any); ok {
				hasFirstName := first["FirstName"]
				fmt.Println(hasFirstName)
				assert.Equal(t, "Lannie", hasFirstName)
			} else {
				t.Errorf("expected first element to be a map, got %T", data[0])
			}
		} else {
			t.Errorf("data array is empty")
		}
	} else {
		t.Errorf("expected data to be a slice, got %T", body["data"])
	}
}
