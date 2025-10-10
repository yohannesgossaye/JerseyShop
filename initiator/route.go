package initiator

import (
	"net/http"

	customerrouting "github.com/yohannesgossaye/internal/glue/routing/customeraccount"
	appmw "github.com/yohannesgossaye/internal/handlers/middleware"
	customerhandler "github.com/yohannesgossaye/internal/handlers/rest/http/customeraccount"

	"github.com/go-chi/chi/v5"
)

// InitRouter registers routes on the provided router and adds middleware
func InitRouter(router chi.Router, handler *customerhandler.CustomerAccountHandler) chi.Router {
	// Middleware to set JSON content type
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	// Register user routes on the SAME router
	customerrouting.RegisterRoutes(router, handler)

	router.Route("/protected", func(r chi.Router) {
		r.Use(appmw.AuthMiddleware)

	})

	return router
}

// InitServer creates the HTTP server with the provided router
func InitServer(router chi.Router) *http.Server {
	port := getEnv("PORT", "8080")

	return &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
}
