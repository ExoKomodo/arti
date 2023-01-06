package artifacts

import (
	"arti/core/operations"
	"arti/lib/api"
)

func (operation UpdateOperation) GetId() operations.OperationId {
	return operation.Id
}

func (operation UpdateOperation) GetArtifact() Artifact {
	return operation.Artifact
}

func NewUpdateOperation(artifact Artifact) (*UpdateOperation, *api.ArtiError) {
	id, err := operations.NewOperationId()
	if err != nil {
		return nil, err
	}
	return &UpdateOperation{
		Id:       operations.OperationId(id),
		Artifact: artifact,
	}, nil
}
