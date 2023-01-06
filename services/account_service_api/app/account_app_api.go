package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joaocarvalhodb1/arch-poc/shared/contracts/protog"
	"github.com/joaocarvalhodb1/arch-poc/shared/helpers"
)

type AccountAppAPI struct {
	httpServer         *http.Server
	accountgRPCService protog.AccountServiceClient
	log                *helpers.Loggers
}

func NewAccountAppAPI(accountService protog.AccountServiceClient, log *helpers.Loggers) *AccountAppAPI {
	app := &AccountAppAPI{
		accountgRPCService: accountService,
		log:                log,
	}
	return app
}

func (app *AccountAppAPI) Routes() http.Handler {
	mux := chi.NewRouter()
	// specify who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	// routes
	mux.Use(middleware.Heartbeat("/ping"))
	mux.Get("/home", app.Home)
	mux.Get("/accounts", app.FindAll)
	mux.Get("/accounts/{id}", app.FindOne)
	mux.Post("/accounts", app.CreateAccount)
	mux.Put("/accounts/{id}", app.UpdateAccount)
	mux.Put("/accounts/credit-limite-apply", app.ApplyCreditLilite)
	mux.Delete("/accounts/{id}", app.DeleteAccount)
	return mux
}
