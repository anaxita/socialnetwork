package main

import (
	"context"
	"github.com/gocraft/dbr"
	"log"
	"net/http"
	"os"
	"os/signal"
	"synergycommunity/internal/bootstrap"
	"synergycommunity/internal/delivery/api"
	"synergycommunity/internal/delivery/api/middleware"
	"synergycommunity/internal/delivery/gql"
	"synergycommunity/internal/domain/interactor"
	"synergycommunity/internal/domain/service"
	"synergycommunity/internal/infrastructure/repository"
	"syscall"
	"time"

	"github.com/pkg/errors"

	_ "github.com/lib/pq"
)

const shutdownTimeout = 5 * time.Second

func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	c, err := bootstrap.NewConfig()
	if err != nil {
		log.Fatalln("Error loading config:", err)
	}

	dbPool, err := bootstrap.NewDBConn(c.DBScheme, c.DBUsername, c.DBPassword, c.DBName, c.DBHost,
		c.DBPort)
	if err != nil {
		log.Fatalln("Error connecting to DB:", err)
	}

	defer func(dbPool *dbr.Connection) {
		err := dbPool.Close()
		if err != nil {
			log.Println("Error closing DB pool:", err)
		}
	}(dbPool)

	repo := repository.NewRepository(dbPool)
	services := service.NewServices(repo)
	interactors := interactor.NewInteractors(services)
	middlewares := middleware.NewMiddlewares(interactors)

	gqlHandler := gql.NewGQLHandler(interactors)
	s := api.NewServer(
		c.HTTPPort,
		gqlHandler,
		middlewares,
	)

	go func() {
		err := s.Start()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalln(err)
		}
	}()

	<-done

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	err = s.Shutdown(ctx)
	if err != nil {
		log.Printf("Error on server shutdown:", err)
	}
}
