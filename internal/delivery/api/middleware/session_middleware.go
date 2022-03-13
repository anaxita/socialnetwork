package middleware

import (
	"context"
	"net/http"
	"synergycommunity/internal/domain/entity"
	"synergycommunity/internal/dto"

	"github.com/google/uuid"
)

type Session struct{}

func NewSession() *Session {
	return &Session{}
}

// Handler creates a session.
func (s *Session) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			session := entity.Session{
				ID: uuid.NewString(),
			}

			r = r.WithContext(context.WithValue(r.Context(), dto.CtxSession, session))

			next.ServeHTTP(w, r)
		},
	)
}
