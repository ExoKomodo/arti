package artifacts

import (
	"arti/core"
	"arti/core/operations"
)

func (operation UpdateOperation) GetId() operations.OperationId {
	return operation.Id
}

func (operation UpdateOperation) GetArtifact() Artifact {
	return operation.Artifact
}

func NewUpdateOperation(artifact Artifact) (*UpdateOperation, *core.ArtiError) {
	id, err := operations.NewOperationId()
	if err != nil {
		return nil, err
	}
	return &UpdateOperation{
		Id:       operations.OperationId(id),
		Artifact: artifact,
	}, nil
}
