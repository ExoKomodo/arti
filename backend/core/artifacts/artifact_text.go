package artifacts

import "arti/core"

func NewTextArtifact(path ArtifactPath, data []byte) (*TextArtifact, *core.ArtiError) {
	return &TextArtifact{
		Data:       data,
		IsArchived: false,
		Path:       path,
	}, nil
}

func (artifact TextArtifact) GetData() []byte {
	return artifact.Data
}

func (artifact TextArtifact) GetIsArchived() bool {
	return artifact.IsArchived
}

func (artifact TextArtifact) GetKind() ArtifactKind {
	return TextArtifactKind
}

func (artifact TextArtifact) GetPath() ArtifactPath {
	return artifact.Path
}
