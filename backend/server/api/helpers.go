package api

import (
	"arti/lib/api"
	"arti/lib/functional"
	"encoding/json"
	"net/http"
)

func marshalResult[T api.Result](w http.ResponseWriter, result T) *api.ArtiError {
	if result.Error() != nil {
		w.WriteHeader(result.Error().HttpCode())
		return result.Error()
	}
	data, err := json.Marshal(result)
	if err != nil {
		inner_err := api.NewArtiError(api.JsonMarshalFailure, err)
		w.WriteHeader(inner_err.HttpCode())
		return inner_err
	}
	w.Write(data)
	return nil
}

func marshalManyResults[T api.Result](w http.ResponseWriter, results []T) *api.ArtiError {
	errors := functional.Map(
		results,
		func(result T) *api.ArtiError {
			return result.Error()
		},
	)
	if functional.Any(
		errors,
		func(err *api.ArtiError) bool {
			return err != nil
		},
	) {
		var wrapped_err *api.ArtiError
		for _, err := range errors {
			wrapped_err = api.NewArtiErrorWrapped(api.MultiRequestFailure, err, wrapped_err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return wrapped_err
	}

	data, err := json.Marshal(results)
	if err != nil {
		inner_err := api.NewArtiError(api.JsonMarshalFailure, err)
		w.WriteHeader(inner_err.HttpCode())
		return inner_err
	}
	w.Write(data)
	return nil
}
