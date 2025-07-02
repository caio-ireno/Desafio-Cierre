package repository

import (
	"app/internal"
	"context"
	"database/sql"
)

func NewRepositoryTicket(db *sql.DB) internal.RepositoryTicket {
	return &mysqlRepository{
		db: db,
	}
}

type mysqlRepository struct {
	db *sql.DB
}

func (r mysqlRepository) GetAll(ctx context.Context) (t map[int]internal.Ticket, err error) {
	println("GetAll chamado no repository MySQL!")
	t = make(map[int]internal.Ticket)
	t[1] = internal.Ticket{
		Id: 1,
		Attributes: internal.TicketAttributes{
			Name:    "Teste",
			Email:   "teste@email.com",
			Country: "Brasil",
			Hour:    "10:00",
			Price:   100.0,
		},
	}
	return
}

func (r *mysqlRepository) GetTicketByDestinationCountry(ctx context.Context, country string) (t map[int]internal.TicketAttributes, err error) {

	return
}

func (r *mysqlRepository) GetTotalAmountTickets(ctx context.Context) (total int, err error) {

	return
}

func (r *mysqlRepository) Update(ctx context.Context, ticket internal.TicketAttributesPatch, id int) (ticketUpdate internal.Ticket, err error) {

	return
}

func (r *mysqlRepository) Create(ctx context.Context, ticket internal.TicketAttributes) (ticketCreated internal.Ticket, err error) {

	return
}
