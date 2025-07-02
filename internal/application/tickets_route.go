package application

import (
	"app/db"
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type router struct {
}

func (router *router) TicketsRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	db := db.GetConnection()
	repository := repository.NewRepositoryTicket(db)
	service := service.NewServiceTicketDefault(repository)
	handler := handler.NewHandlerTicketDefault(service)

	r.Route("/tickets", func(r chi.Router) {
		r.Get("/", handler.GetAll())
		r.Post("/", handler.Create())
		r.Get("/total_amount", handler.GetTotalAmountTickets())
		r.Patch("/update/{id}", handler.Update())
	})

	return r
}

func NewRouter() *router {
	return &router{}
}
