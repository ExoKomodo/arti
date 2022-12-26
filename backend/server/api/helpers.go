package api

import (
	"encoding/json"
	"net/http"
)

func marshalResult[T any](w http.ResponseWriter, results T) error {
	data, err := json.Marshal(results)
	if err != nil {
		return err
	}
	w.Write(data)
	return nil
}
