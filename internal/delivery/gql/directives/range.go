package directives

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"synergycommunity/internal/domain"
)

func Range(ctx context.Context, _ interface{}, next graphql.Resolver, min int64, max int64) (interface{}, error) {
	val, err := next(ctx)
	if err != nil {
		return nil, domain.NewError(domain.ErrCodeInternal)
	}

	var l int

	switch typedVal := val.(type) {
	case string:
		l = len(typedVal)
	case []int64:
		l = len(typedVal)
	case []string:
		l = len(typedVal)
	default:
		return nil, domain.NewError(domain.ErrCodeBadRequest, val, "must be a string or array")
	}

	if l < int(min) || l > int(max) {
		return nil, domain.NewError(domain.ErrCodeValidationErr)
	}

	return val, nil
}
