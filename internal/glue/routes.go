package glue

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Route struct {
	Method      string
	Path        string
	Handler     http.HandlerFunc
	Middlewares []func(next http.Handler) http.Handler
}

func RegisterRoutes(router chi.Router, routes []Route) {
	for _, route := range routes {
		// wrap handler with any per-route middlewares (outermost last)
		var h http.Handler = route.Handler
		if len(route.Middlewares) > 0 {
			for i := len(route.Middlewares) - 1; i >= 0; i-- {
				h = route.Middlewares[i](h)
			}
		}

		switch route.Method {
		case http.MethodGet:
			router.Get(route.Path, h.ServeHTTP)
		case http.MethodPost:
			router.Post(route.Path, h.ServeHTTP)
		case http.MethodPatch:
			router.Patch(route.Path, h.ServeHTTP)
		case http.MethodDelete:
			router.Delete(route.Path, h.ServeHTTP)
		default:
			router.Method(route.Method, route.Path, h)
		}
	}
}
