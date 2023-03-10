package artifacts

import "arti/core/operations"

func (operation ArchiveOperation) GetId() operations.OperationId {
	return operation.Id
}

func (operation ArchiveOperation) GetPath() ArtifactPath {
	return operation.Path
}

func NewArchiveOperation(path ArtifactPath) (*ArchiveOperation, error) {
	id, err := operations.NewOperationId()
	if err != nil {
		return nil, err
	}
	return &ArchiveOperation{
		Id:   operations.OperationId(id),
		Path: path,
	}, nil
}
