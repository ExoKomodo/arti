package api

import (
	"fmt"
	"net/http"
)

type ErrorKind = string
type ArtiError struct {
	Kind       ErrorKind
	Err        error
	WrappedErr error
}

func NewArtiError(kind ErrorKind, err error) *ArtiError {
	return &ArtiError{
		Kind:       kind,
		Err:        err,
		WrappedErr: nil,
	}
}

func NewArtiErrorWrapped(kind ErrorKind, err error, wrappedErr error) *ArtiError {
	return &ArtiError{
		Kind:       kind,
		Err:        err,
		WrappedErr: wrappedErr,
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
	if err.WrappedErr != nil {
		return fmt.Sprintf("%s:%d:%v:%v", err.Kind, err.HttpCode(), err.Err, err.WrappedErr)
	}
	return fmt.Sprintf("%s:%d:%v", err.Kind, err.HttpCode(), err.Err)
}
