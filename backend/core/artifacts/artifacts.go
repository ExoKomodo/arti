package artifacts

import (
	"arti/core"
	"fmt"
)

func NewArtifact(kind ArtifactKind, path ArtifactPath, data []byte) (Artifact, *core.ArtiError) {
	switch kind {
	case TextArtifactKind:
		return NewTextArtifact(path, data)
	}
	return nil, core.NewArtiError(core.UnsupportedArtifactKind, fmt.Errorf("%s", kind))
}
