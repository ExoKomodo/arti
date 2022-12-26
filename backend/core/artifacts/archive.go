package artifacts

import "arti/lib/functional"

func Archive(request *ArchiveRequest) *ArchiveResult {
	return ArchiveByPath(request.Path)
}

func ArchiveMany(requests []*ArchiveRequest) []*ArchiveResult {
	return functional.GoMap(requests, Archive)
}

func ArchiveByPath(path ArtifactPath) *ArchiveResult {
	operation, err := NewArchiveOperation(path)
	return &ArchiveResult{
		Operation: operation,
		Err:       err,
	}
}

func ArchiveManyByPath(paths []ArtifactPath) []*ArchiveResult {
	return functional.GoMap(paths, ArchiveByPath)
}
