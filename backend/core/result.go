package core

type Result interface {
	Error() *ArtiError
}
