package customeraccount

import (
	"github.com/go-chi/chi/v5"
	customerport "github.com/yohannesgossaye/internal/domain/interfaces/customeraccount"
	"github.com/yohannesgossaye/internal/glue"
	"net/http"
)

func RegisterRoutes(router chi.Router, customerAccountHandler customerport.CustomerAccountHandler) {
	routes := []glue.Route{
		{
			Method:  http.MethodPost,
			Path:    "/customer",
			Handler: customerAccountHandler.CreateCustomer,
		},
		{
			Method:  http.MethodPost,
			Path:    "/customer/login",
			Handler: customerAccountHandler.LoginCustomer,
		},
		{
			Method:  http.MethodPost,
			Path:    "/customer/verify-otp",
			Handler: customerAccountHandler.VerifyCustomerOtp,
		},
	}
	glue.RegisterRoutes(router, routes)
}
