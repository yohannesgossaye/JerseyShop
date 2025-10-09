package customeraccount

import (
	"encoding/json"
	"net/http"

	"github.com/yohannesgossaye/internal/domain/dto/customeraccount"
	core "github.com/yohannesgossaye/internal/handlers/rest/http/customeraccount/core"
	CustomerService "github.com/yohannesgossaye/internal/service/customeraccount"
)

type CustomerAccountHandler struct {
	service CustomerService.CustomerService
}

func NewCustomerAccountHandler(service CustomerService.CustomerService) *CustomerAccountHandler {
	return &CustomerAccountHandler{
		service: service,
	}
}

func (h *CustomerAccountHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var req customeraccount.CreateCustomer
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := req.Validator(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	CreatedMap := core.MapCreateCustomerRequestToModel(req)
	customer, err := h.service.CreateCustomer(r.Context(), CreatedMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}

func (h *CustomerAccountHandler) LoginCustomer(w http.ResponseWriter, r *http.Request) {
	var req customeraccount.LoginCustomer
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := req.Validator(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	customer, err := h.service.LoginCustomer(r.Context(), req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}

func (h *CustomerAccountHandler) VerifyCustomerOtp(w http.ResponseWriter, r *http.Request) {
	var req customeraccount.VerifyCustomerOtp
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := req.Validator(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	customer, err := h.service.VerifyCustomerOtp(r.Context(), req.Email, req.OTPCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}
