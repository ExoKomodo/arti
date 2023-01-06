package api

const (
	JsonMarshalFailure         ErrorKind = "json_marshal_failure"
	MultiRequestFailure        ErrorKind = "multi_request_failure"
	OperationIdCreationFailure ErrorKind = "operation_id_creation_failure"
	UnknownErrorKind           ErrorKind = "unknown_error_kind"
	UnsupportedArtifactKind    ErrorKind = "unsupported_artifact_kind"
)

var ArtiErrorKinds []ErrorKind = []ErrorKind{
	JsonMarshalFailure,
	MultiRequestFailure,
	OperationIdCreationFailure,
	UnknownErrorKind,
	UnsupportedArtifactKind,
}
