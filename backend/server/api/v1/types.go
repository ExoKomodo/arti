package v1

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ApiMethod struct {
	Route   ApiRoute
	Handler func(http.ResponseWriter, *http.Request)
	Method  string
}

type ApiRoute = string

type Controller struct {
	Methods map[ApiRoute]ApiMethod
}

func (controller Controller) Register(router *chi.Mux) error {
	for route, method := range controller.Methods {
		switch method.Method {
		case http.MethodDelete:
			router.Delete(route, method.Handler)
		case http.MethodGet:
			router.Get(route, method.Handler)
		case http.MethodPost:
			router.Post(route, method.Handler)
		case http.MethodPut:
			router.Put(route, method.Handler)
		default:
			return fmt.Errorf("failed to register route: %s %s", method.Method, route)
		}
	}
	return nil
}
