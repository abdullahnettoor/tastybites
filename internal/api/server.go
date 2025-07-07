package api

import (
	"net/http"
	"strconv"

	"github.com/abdullahnettoor/tastybites/internal/api/middlewares"
	"github.com/abdullahnettoor/tastybites/internal/config"
	"github.com/abdullahnettoor/tastybites/internal/repo/interfaces"
)

type application struct {
	Config *config.Config
	Repo   interfaces.Repository
	Server *http.Server
	Mux    *http.ServeMux
}

func NewApp(config *config.Config, repo interfaces.Repository) (*application, error) {
	mux := http.NewServeMux()

	app := &application{
		Config: config,
		Repo:   repo,
		Mux:    mux,
	}

	handler := middlewares.MiddlewareChain(
		mux,
		middlewares.RecoverPanic,
		middlewares.CORS,
		middlewares.LogReq,
	)

	app.Server = &http.Server{
		Addr:    config.ServerConfig.Host + ":" + strconv.Itoa(config.ServerConfig.Port),
		Handler: handler,
	}

	return app, nil
}

func (app *application) Start() error {
	if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
