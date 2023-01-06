package artifacts

import (
	"arti/core/operations"
	"arti/lib/api"
)

func (operation DeleteOperation) GetId() operations.OperationId {
	return operation.Id
}

func (operation DeleteOperation) GetPath() ArtifactPath {
	return operation.Path
}

func NewDeleteOperation(path ArtifactPath) (*DeleteOperation, *api.ArtiError) {
	id, err := operations.NewOperationId()
	if err != nil {
		return nil, err
	}
	return &DeleteOperation{
		Id:   operations.OperationId(id),
		Path: path,
	}, nil
}
