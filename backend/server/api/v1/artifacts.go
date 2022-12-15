package v1

import (
	"arti/lib/artifacts"
	"net/http"
)

func getArtifacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.Write([]byte(artifacts.GetAll()))
}

var (
	GetArtifacts ApiMethod = ApiMethod{
		Route:   "/v1/artifacts",
		Handler: getArtifacts,
		Method:  http.MethodGet,
	}
)

func NewArtifactsController() *Controller {
	return &Controller{
		Methods: map[ApiRoute]ApiMethod{
			GetArtifacts.Route: GetArtifacts,
		},
	}
}
