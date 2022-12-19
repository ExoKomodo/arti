package api

import (
	"net/http"
)

var (
	GetStatus ApiMethod = ApiMethod{
		Route:   "/status",
		Handler: getStatus,
		Method:  http.MethodGet,
	}
)

func NewRootController() Controller {
	return &RootController{}
}

func (controller RootController) BaseUrl() string {
	return "/"
}

func (controller RootController) Routes() []*ApiMethod {
	return []*ApiMethod{
		&GetStatus,
	}
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte("Server is running!"))
}
