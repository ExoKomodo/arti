package artifacts

import "fmt"

func NewArtifact(kind ArtifactKind, path ArtifactPath, data []byte) (Artifact, error) {
	switch kind {
	case TextArtifactKind:
		return NewTextArtifact(path, data)
	}
	return nil, fmt.Errorf("unsupported artifact kind: %s", kind)
}
