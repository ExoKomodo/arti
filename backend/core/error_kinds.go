package core

const (
	IdCreationFailure       ErrorKind = "id_creation_failure"
	JsonMarshalFailure      ErrorKind = "json_marshal_failure"
	UnknownErrorKind        ErrorKind = "unknown_error_kind"
	UnsupportedArtifactKind ErrorKind = "unsupported_artifact_kind"
)

var ArtiErrorKinds []ErrorKind = []ErrorKind{
	IdCreationFailure,
	JsonMarshalFailure,
	UnknownErrorKind,
	UnsupportedArtifactKind,
}
