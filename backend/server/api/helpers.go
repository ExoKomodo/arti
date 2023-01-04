package api

import (
	"arti/core"
	"encoding/json"
	"net/http"
)

func marshalResult[T core.Result](w http.ResponseWriter, result T) *core.ArtiError {
	if result.Error() != nil {
		w.WriteHeader(result.Error().HttpCode())
		return result.Error()
	}
	data, err := json.Marshal(result)
	if err != nil {
		wrapped_err := core.NewArtiError(core.JsonMarshalFailure, err)
		w.WriteHeader(wrapped_err.HttpCode())
		return wrapped_err
	}
	w.Write(data)
	return nil
}

func marshalManyResults[T []core.Result](w http.ResponseWriter, results T) *core.ArtiError {
	data, err := json.Marshal(results)
	if err != nil {
		wrapped_err := core.NewArtiError(core.JsonMarshalFailure, err)
		w.WriteHeader(wrapped_err.HttpCode())
		return wrapped_err
	}
	w.Write(data)
	return nil
}
