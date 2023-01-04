package core

import (
	"fmt"
	"net/http"
)

type ErrorKind = string
type ArtiError struct {
	Kind ErrorKind
	Err  error
}

func NewArtiError(kind ErrorKind, err error) *ArtiError {
	return &ArtiError{
		Kind: kind,
		Err:  err,
	}
}

func (info *ArtiError) HttpCode() int {
	switch info.Kind {
	case UnsupportedArtifactKind:
		return http.StatusUnsupportedMediaType
	case UnknownErrorKind:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func (err *ArtiError) Error() string {
	return fmt.Sprintf("%s:%d:%v", err.Kind, err.HttpCode(), err.Err)
}
