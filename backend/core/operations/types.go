package operations

type OperationId string

type Operation interface {
	GetId() OperationId
}
