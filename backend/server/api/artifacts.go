package api

import (
	"arti/core/artifacts"
	"fmt"
	"net/http"
)

var (
	ArchiveArtifact ApiMethod = ApiMethod{
		Route:   "/archive/{path}",
		Handler: archiveArtifact,
		Method:  http.MethodDelete,
	}
	ArchiveManyArtifacts ApiMethod = ApiMethod{
		Route:   "/archive/many",
		Handler: archiveManyArtifacts,
		Method:  http.MethodDelete,
	}
	DeleteArtifact ApiMethod = ApiMethod{
		Route:   "/{path}",
		Handler: deleteArtifact,
		Method:  http.MethodDelete,
	}
	DeleteManyArtifacts ApiMethod = ApiMethod{
		Route:   "/many",
		Handler: deleteManyArtifacts,
		Method:  http.MethodDelete,
	}
	GetArtifact ApiMethod = ApiMethod{
		Route:   "/{path}",
		Handler: getArtifact,
		Method:  http.MethodGet,
	}
	GetManyArtifacts ApiMethod = ApiMethod{
		Route:   "/",
		Handler: getManyArtifacts,
		Method:  http.MethodGet,
	}
	UpdateArtifact ApiMethod = ApiMethod{
		Route:   "/",
		Handler: updateArtifact,
		Method:  http.MethodPatch,
	}
	UpdateManyArtifacts ApiMethod = ApiMethod{
		Route:   "/many",
		Handler: updateManyArtifacts,
		Method:  http.MethodPatch,
	}
	UploadArtifact ApiMethod = ApiMethod{
		Route:   "/",
		Handler: uploadArtifact,
		Method:  http.MethodPost,
	}
	UploadManyArtifacts ApiMethod = ApiMethod{
		Route:   "/many",
		Handler: uploadManyArtifacts,
		Method:  http.MethodPost,
	}
)

func NewArtifactsController() Controller {
	return ArtifactsController{}
}

func (controller ArtifactsController) BaseUrl() string {
	return "/artifacts"
}

func (controller ArtifactsController) Routes() []*ApiMethod {
	return []*ApiMethod{
		&ArchiveArtifact,
		&ArchiveManyArtifacts,

		&DeleteArtifact,
		&DeleteManyArtifacts,

		&GetArtifact,
		&GetManyArtifacts,

		&UpdateArtifact,
		&UpdateManyArtifacts,

		&UploadArtifact,
		&UploadManyArtifacts,
	}
}

func archiveArtifact(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/json")
	results := artifacts.ArchiveByPath("path")
	marshalResult(w, results)
}

func archiveManyArtifacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/json")
	paths := []artifacts.ArtifactPath{"test"}
	results := artifacts.ArchiveManyByPath(paths)
	marshalResult(w, results)
}

func deleteArtifact(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/json")
	results := artifacts.DeleteByPath("path")
	marshalResult(w, results)
}

func deleteManyArtifacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/json")
	paths := []artifacts.ArtifactPath{"test"}
	results := artifacts.DeleteManyByPath(paths)
	marshalResult(w, results)
}

func getArtifact(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/json")
	path := artifacts.ArtifactPath("test")
	results := artifacts.GetByPath(path)
	marshalResult(w, results)
}

func getManyArtifacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/json")
	paths := []artifacts.ArtifactPath{artifacts.ArtifactPath("test")}
	results := artifacts.GetManyByPath(paths)
	marshalResult(w, results)
}

func updateArtifact(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/json")
	artifact, err := artifacts.NewTextArtifact(artifacts.ArtifactPath("test"), []byte{})
	if err != nil {
		fmt.Println("failed to update artifact")
		return
	}
	results := artifacts.Update(
		&artifacts.UpdateRequest{
			Artifact: artifact,
		},
	)
	marshalResult(w, results)
}

func updateManyArtifacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/json")
	artifact, err := artifacts.NewTextArtifact(artifacts.ArtifactPath("test"), []byte{})
	if err != nil {
		fmt.Println("failed to update many artifacts")
		return
	}
	results := artifacts.UpdateMany(
		[]*artifacts.UpdateRequest{
			{
				Artifact: artifact,
			},
		},
	)
	marshalResult(w, results)
}

func uploadArtifact(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/json")
	artifact, err := artifacts.NewTextArtifact(artifacts.ArtifactPath("test"), []byte{})
	if err != nil {
		fmt.Println("failed to upload text artifact")
		return
	}
	results := artifacts.Upload(
		&artifacts.UploadRequest{
			Artifact: artifact,
		},
	)
	marshalResult(w, results)
}

func uploadManyArtifacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/json")
	artifact, err := artifacts.NewTextArtifact(artifacts.ArtifactPath("test"), []byte{})
	if err != nil {
		fmt.Println("failed to upload many text artifacts")
		return
	}
	results := artifacts.UploadMany(
		[]*artifacts.UploadRequest{
			{
				Artifact: artifact,
			},
		},
	)
	marshalResult(w, results)
}
