package artifacts

import (
	"arti/core"
	"arti/core/operations"
)

func (operation UploadOperation) GetId() operations.OperationId {
	return operation.Id
}

func (operation UploadOperation) GetArtifact() Artifact {
	return operation.Artifact
}

func NewUploadOperation(artifact Artifact) (*UploadOperation, *core.ArtiError) {
	id, err := operations.NewOperationId()
	if err != nil {
		return nil, err
	}
	return &UploadOperation{
		Id:       operations.OperationId(id),
		Artifact: artifact,
	}, nil
}
