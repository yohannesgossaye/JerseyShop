package customeraccount

import (
	"encoding/json"
	"net/http"

	"github.com/yohannesgossaye/internal/domain/dto/customeraccount"
	core "github.com/yohannesgossaye/internal/handlers/rest/http/customeraccount/core"
	CustomerService "github.com/yohannesgossaye/internal/service/customeraccount"
	response "github.com/yohannesgossaye/pkgs/message/Response"
	"github.com/yohannesgossaye/pkgs/message/localization"

	logger "github.com/yohannesgossaye/pkgs/logger"
)

type CustomerAccountHandler struct {
	service CustomerService.CustomerService
	log     logger.Logger
}

func NewCustomerAccountHandler(service CustomerService.CustomerService, log logger.Logger) *CustomerAccountHandler {
	return &CustomerAccountHandler{
		service: service,
		log:     log,
	}
}

func (h *CustomerAccountHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	h.log.Infof("received request to create customer")
	var req customeraccount.CreateCustomer
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.log.Errorf("Error decoding request body: %v", err)
		localization.SendErrorByCodeResponse(w, localization.ErrCustomerCreateDecodeFailed.Code)
		return
	}
	if err := req.Validator(); err != nil {
		h.log.Errorf("Error validating request: %v", err)
		localization.SendErrorByCodeResponse(w, localization.ErrorValidationFailed.Code)
		return
	}
	CreatedMap := core.MapCreateCustomerRequestToModel(req)
	customer, err := h.service.CreateCustomer(r.Context(), CreatedMap)
	if err != nil {
		h.log.Errorf("Error creating customer: %v", err)
		localization.SendErrorByCodeResponse(w, localization.ErrCustomerCreate.Code)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	h.log.Infof("customer created successfully")
	response.SendSuccessResponse(w, http.StatusCreated, "Customer Registered successfully", customer, nil)
}

func (h *CustomerAccountHandler) LoginCustomer(w http.ResponseWriter, r *http.Request) {
	h.log.Infof("received request to login customer")
	var req customeraccount.LoginCustomer
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.log.Errorf("Error decoding request body: %v", err)
		localization.SendErrorByCodeResponse(w, localization.ErrCustomerCreateDecodeFailed.Code)
		return
	}
	if err := req.Validator(); err != nil {
		h.log.Errorf("Error validating request: %v", err)
		localization.SendErrorByCodeResponse(w, localization.ErrorValidationFailed.Code)
		return
	}
	customer, err := h.service.LoginCustomer(r.Context(), req.Email, req.Password)
	if err != nil {
		h.log.Errorf("Error logging in customer: %v", err)
		localization.SendErrorByCodeResponse(w, localization.ErrCustomerLogin.Code)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h.log.Infof("customer logged in successfully")
	response.SendSuccessResponse(w, http.StatusOK, "customer logged in successfully", customer, nil)
}

func (h *CustomerAccountHandler) VerifyCustomerOtp(w http.ResponseWriter, r *http.Request) {
	h.log.Infof("received request to verify customer otp")
	var req customeraccount.VerifyCustomerOtp
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		h.log.Errorf("Error decoding request body: %v", err)
		localization.SendErrorByCodeResponse(w, localization.ErrCustomerCreateDecodeFailed.Code)
		return
	}
	if err := req.Validator(); err != nil {
		h.log.Errorf("Error validating request: %v", err)
		localization.SendErrorByCodeResponse(w, localization.ErrorValidationFailed.Code)
		return
	}
	customer, err := h.service.VerifyCustomerOtp(r.Context(), req.Email, req.OTPCode)
	if err != nil {
		h.log.Errorf("Error verifying customer otp: %v", err)
		// Map specific messages to specific error codes
		switch err.Error() {
		case localization.ErrInvalidOtp.Message:
			localization.SendErrorByCodeResponse(w, localization.ErrInvalidOtp.Code)
		case localization.ErrOtpExpired.Message:
			localization.SendErrorByCodeResponse(w, localization.ErrOtpExpired.Code)
		default:
			localization.SendErrorByCodeResponse(w, localization.ErrCustomerNotVerified.Code)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	h.log.Infof("customer otp verified successfully")
	response.SendSuccessResponse(w, http.StatusOK, "customer otp verified successfully", customer, nil)
}
