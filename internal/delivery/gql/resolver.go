package gql

//go:generate go run github.com/99designs/gqlgen generate

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-playground/validator/v10"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"log"
	"net/http"
	"synergycommunity/internal/delivery/api/helpers"
	"synergycommunity/internal/delivery/gql/directives"
	"synergycommunity/internal/domain/interactor"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	v *validator.Validate
	*interactor.Interactors
}

func NewGQLHandler(i *interactor.Interactors) http.Handler {
	srv := handler.NewDefaultServer(
		NewExecutableSchema(
			Config{
				Resolvers: &Resolver{
					validator.New(),
					i,
				},
				Directives: DirectiveRoot{
					Range: directives.Range,
				},
			},
		),
	)

	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		err := graphql.DefaultErrorPresenter(ctx, e)

		myErr := helpers.SetCtxToError(ctx, e)
		err.Extensions = myErr.ToMap()

		log.Println(e)

		return err
	})

	return srv
}
