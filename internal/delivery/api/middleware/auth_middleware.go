package middleware

import (
	"context"
	"net/http"
	"strconv"
	"synergycommunity/internal/domain/interactor"
	"synergycommunity/internal/dto"
)

type Auth struct {
	users *interactor.UserInteractor
}

func NewAuth(users *interactor.UserInteractor) *Auth {
	return &Auth{users: users}
}

// Handler creates a new callback that is run when we require authentication.
func (s *Auth) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			var (
				ctx     = r.Context()
				session = dto.Session(ctx)
			)

			username, password, isAuth := r.BasicAuth()
			if !isAuth {
				next.ServeHTTP(w, r)

				return
			}

			// в будущем из токена мы будем по айди искать поэтому считаем что нам айди пришёл
			userID, err := strconv.ParseInt(username, 10, 64)
			if err != nil {
				next.ServeHTTP(w, r)

				return
			}

			user, err := s.users.UserByID(ctx, userID)
			if err != nil {
				next.ServeHTTP(w, r)

				return
			}

			// чтобы не добавлять поле password во все модели, просто захардкодим пароль
			if password != "dev125" {
				next.ServeHTTP(w, r)

				return
			}

			session.User = user

			r = r.WithContext(context.WithValue(r.Context(), dto.CtxSession, session))

			next.ServeHTTP(w, r)
		},
	)
}
