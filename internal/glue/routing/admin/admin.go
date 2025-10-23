package admin

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	adminport "github.com/yohannesgossaye/internal/domain/interfaces/admin"
	"github.com/yohannesgossaye/internal/glue"
	appmw "github.com/yohannesgossaye/internal/handlers/middleware"
)

func RegisterRoutes(router chi.Router, adminAccountHandler adminport.AdminAccountHandler) {
	routes := []glue.Route{
		{
			Method:  http.MethodPost,
			Path:    "/admin",
			Handler: adminAccountHandler.CreateAdminAccount,
		},
		{
			Method:  http.MethodPost,
			Path:    "/admin/login",
			Handler: adminAccountHandler.LoginAdminAccount,
			Middlewares: []func(next http.Handler) http.Handler{
				appmw.JWTAuthMiddleware,
			},
		},
		{
			Method:  http.MethodGet,
			Path:    "/admin/{id}",
			Handler: adminAccountHandler.GetAdminAccount,
		},
		{
			Method:  http.MethodGet,
			Path:    "/admin/all",
			Handler: adminAccountHandler.GetAdminAccounts,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/admin",
			Handler: adminAccountHandler.DeleteAdminAccount,
		},
	}
	glue.RegisterRoutes(router, routes)
}
