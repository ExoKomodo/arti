package operations

import (
	"arti/core"

	"github.com/google/uuid"
)

func NewOperationId() (OperationId, *core.ArtiError) {
	id, err := uuid.NewRandom()
	if err != nil {
		// Returns 00000000-0000-0000-0000-000000000000
		return OperationId(uuid.NullUUID{}.UUID.String()), core.NewArtiError(core.IdCreationFailure, err)
	}
	return OperationId(id.String()), nil
}
