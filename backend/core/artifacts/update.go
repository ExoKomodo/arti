package artifacts

import "arti/lib/functional"

func Update(request *UpdateRequest) *UpdateResult {
	operation, err := NewUpdateOperation(request.Artifact)
	return &UpdateResult{
		Operation: operation,
		Err:       err,
	}
}

func UpdateMany(requests []*UpdateRequest) []*UpdateResult {
	return functional.GoMap(requests, Update)
}
