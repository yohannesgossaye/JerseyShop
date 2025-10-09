package glue

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func RegisterRoutes(router chi.Router, routes []Route) {
	for _, route := range routes {
		switch route.Method {
		case http.MethodGet:
			router.Get(route.Path, route.Handler)
		case http.MethodPost:
			router.Post(route.Path, route.Handler)
		case http.MethodPatch:
			router.Patch(route.Path, route.Handler)
		case http.MethodDelete:
			router.Delete(route.Path, route.Handler)
		default:
			router.Method(route.Method, route.Path, route.Handler)
		}
	}
}
