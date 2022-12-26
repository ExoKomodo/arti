package artifacts

import "arti/lib/functional"

func Get(request *GetRequest) *GetResult {
	artifact, err := NewTextArtifact(request.Path, []byte{})
	return &GetResult{
		Artifact: artifact,
		Err:      err,
	}
}

func GetMany(requests []*GetRequest) []*GetResult {
	return functional.GoMap(requests, Get)
}

func GetByPath(path ArtifactPath) *GetResult {
	artifact, err := NewTextArtifact(path, []byte{})
	return &GetResult{
		Artifact: artifact,
		Err:      err,
	}
}

func GetManyByPath(paths []ArtifactPath) []*GetResult {
	return functional.GoMap(paths, GetByPath)
}
