package artifacts

import "arti/lib/functional"

func Delete(request *DeleteRequest) *DeleteResult {
	return DeleteByPath(request.Path)
}

func DeleteMany(requests []*DeleteRequest) []*DeleteResult {
	return functional.GoMap(requests, Delete)
}

func DeleteByPath(path ArtifactPath) *DeleteResult {
	operation, err := NewDeleteOperation(path)
	return &DeleteResult{
		Operation: operation,
		Err:       err,
	}
}

func DeleteManyByPath(paths []ArtifactPath) []*DeleteResult {
	return functional.GoMap(paths, DeleteByPath)
}
