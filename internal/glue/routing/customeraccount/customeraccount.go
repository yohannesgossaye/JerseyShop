package customeraccount

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	customerport "github.com/yohannesgossaye/internal/domain/interfaces/customeraccount"
	"github.com/yohannesgossaye/internal/glue"
	appmw "github.com/yohannesgossaye/internal/handlers/middleware"
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
			Middlewares: []func(next http.Handler) http.Handler{
				appmw.AuthMiddleware,
			},
		},
		{
			Method:  http.MethodPost,
			Path:    "/customer/verify-otp",
			Handler: customerAccountHandler.VerifyCustomerOtp,
		},
	}
	glue.RegisterRoutes(router, routes)
}
