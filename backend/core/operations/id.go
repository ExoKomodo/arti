package operations

import (
	"arti/lib/api"

	"github.com/google/uuid"
)

func NewOperationId() (OperationId, *api.ArtiError) {
	id, err := uuid.NewRandom()
	if err != nil {
		// Returns 00000000-0000-0000-0000-000000000000
		return OperationId(uuid.NullUUID{}.UUID.String()), api.NewArtiError(api.OperationIdCreationFailure, err)
	}
	return OperationId(id.String()), nil
}
