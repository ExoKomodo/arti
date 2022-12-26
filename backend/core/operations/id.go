package operations

import "github.com/google/uuid"

func NewOperationId() (OperationId, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		// Returns 00000000-0000-0000-0000-000000000000
		return OperationId(uuid.NullUUID{}.UUID.String()), err
	}
	return OperationId(id.String()), err
}
