package dto

import (
	"context"
	"synergycommunity/internal/domain/entity"
)

type ctxKeySession string

const CtxSession ctxKeySession = "session"

func Session(ctx context.Context) entity.Session {
	value := ctx.Value(CtxSession)

	session, ok := value.(entity.Session)
	if !ok {
		return entity.Session{}
	}

	return session
}
