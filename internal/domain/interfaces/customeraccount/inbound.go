package customeraccount

import (
	"net/http"
)

type CustomerAccountHandler interface {
	CreateCustomer(w http.ResponseWriter, r *http.Request)
	LoginCustomer(w http.ResponseWriter, r *http.Request)
	VerifyCustomerOtp(w http.ResponseWriter, r *http.Request)
}
