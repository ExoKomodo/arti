package artifacts

import "arti/core/operations"

func (operation DeleteOperation) GetId() operations.OperationId {
	return operation.Id
}

func (operation DeleteOperation) GetPath() ArtifactPath {
	return operation.Path
}

func NewDeleteOperation(path ArtifactPath) (*DeleteOperation, error) {
	id, err := operations.NewOperationId()
	if err != nil {
		return nil, err
	}
	return &DeleteOperation{
		Id:   operations.OperationId(id),
		Path: path,
	}, nil
}
