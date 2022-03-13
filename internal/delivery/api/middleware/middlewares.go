package middleware

import "synergycommunity/internal/domain/interactor"

// M provides list of all middlewares.
type M struct {
	Auth    *Auth
	Cors    *Cors
	Session *Session
}

func NewMiddlewares(i *interactor.Interactors) *M {
	return &M{
		Auth:    NewAuth(i.Users),
		Cors:    NewCors(),
		Session: NewSession(),
	}
}
