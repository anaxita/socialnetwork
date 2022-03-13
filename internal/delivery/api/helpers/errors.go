package helpers

import (
	"context"
	"errors"
	"synergycommunity/internal/domain"
	"synergycommunity/internal/dto"
)

func SetCtxToError(ctx context.Context, e error) *domain.Error {
	if errors.Unwrap(e) == nil {
		e = domain.NewError(domain.ErrCodeInternal, "errors is empty, check code")
	}

	var err *domain.Error

	ok := errors.As(e, &err)
	if !ok {
		err = domain.NewError(domain.ErrCodeInternal, "errors has incorrect type, check code")
	}

	session := dto.Session(ctx)

	err.UUID = session.ID
	err.UserID = session.User.ID

	return err
}
