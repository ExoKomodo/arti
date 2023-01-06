package artifacts

import (
	"arti/core/operations"
	"arti/lib/api"
)

type ArtifactKind string

const Unknown ArtifactKind = ""

type ArtifactPath string

/*************/
/* Artifacts */
/*************/
type Artifact interface {
	GetData() []byte
	GetIsArchived() bool
	GetKind() ArtifactKind
	GetPath() ArtifactPath
}

type TextArtifact struct {
	Data       []byte       `json:"data"`
	IsArchived bool         `json:"isArchived"`
	Path       ArtifactPath `json:"path"`
}

/***********/
/* Archive */
/***********/
type ArchiveOperation struct {
	Id   operations.OperationId `json:"id"`
	Path ArtifactPath           `json:"path"`
}

type ArchiveRequest struct {
	Path ArtifactPath `json:"path"`
}

type ArchiveResult struct {
	Operation *ArchiveOperation `json:"operation"`
	Err       *api.ArtiError    `json:"err"`
}

func (r *ArchiveResult) Error() *api.ArtiError {
	return r.Err
}

/**********/
/* Delete */
/**********/
type DeleteRequest struct {
	Path ArtifactPath `json:"path"`
}

type DeleteOperation struct {
	Id   operations.OperationId `json:"id"`
	Path ArtifactPath           `json:"path"`
}

type DeleteResult struct {
	Operation *DeleteOperation `json:"operation"`
	Err       *api.ArtiError   `json:"err"`
}

func (r *DeleteResult) Error() *api.ArtiError {
	return r.Err
}

/*******/
/* Get */
/*******/
type GetRequest struct {
	Path ArtifactPath `json:"path"`
}

type GetResult struct {
	Artifact Artifact       `json:"artifact"`
	Err      *api.ArtiError `json:"err"`
}

func (r *GetResult) Error() *api.ArtiError {
	return r.Err
}

/**********/
/* Update */
/**********/
type UpdateRequest struct {
	Artifact Artifact `json:"artifact"`
}

type UpdateOperation struct {
	Id       operations.OperationId `json:"id"`
	Artifact Artifact               `json:"artifact"`
}

type UpdateResult struct {
	Operation *UpdateOperation `json:"operation"`
	Err       *api.ArtiError   `json:"err"`
}

func (r *UpdateResult) Error() *api.ArtiError {
	return r.Err
}

/**********/
/* Upload */
/**********/
type UploadRequest struct {
	Artifact Artifact `json:"artifact"`
}

type UploadOperation struct {
	Id       operations.OperationId `json:"id"`
	Artifact Artifact               `json:"artifact"`
}

type UploadResult struct {
	Operation *UploadOperation `json:"operation"`
	Err       *api.ArtiError   `json:"err"`
}

func (r *UploadResult) Error() *api.ArtiError {
	return r.Err
}
