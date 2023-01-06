package artifacts

import (
	"arti/lib/api"
	"fmt"
)

func NewArtifact(kind ArtifactKind, path ArtifactPath, data []byte) (Artifact, *api.ArtiError) {
	switch kind {
	case TextArtifactKind:
		return NewTextArtifact(path, data)
	}
	return nil, api.NewArtiError(api.UnsupportedArtifactKind, fmt.Errorf("%s", kind))
}
