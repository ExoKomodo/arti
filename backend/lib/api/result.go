package api

type Result interface {
	Error() *ArtiError
}
