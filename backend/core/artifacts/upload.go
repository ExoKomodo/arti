package artifacts

import "arti/lib/functional"

func Upload(request *UploadRequest) *UploadResult {
	operation, err := NewUploadOperation(request.Artifact)
	return &UploadResult{
		Operation: operation,
		Err:       err,
	}
}

func UploadMany(requests []*UploadRequest) []*UploadResult {
	return functional.GoMap(requests, Upload)
}
