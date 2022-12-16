package v1

import (
	"net/http"
)

func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte("Server is running!"))
}

var (
	GetStatus ApiMethod = ApiMethod{
		Route:   "/v1/status",
		Handler: getStatus,
		Method:  http.MethodGet,
	}
)

func NewRootController() *Controller {
	return &Controller{
		Methods: map[ApiRoute]ApiMethod{
			GetStatus.Route: GetStatus,
		},
	}
}
