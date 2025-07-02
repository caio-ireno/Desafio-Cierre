package repository

import (
	"app/internal"
	"app/pkg/apperrors"
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
	rows, err := r.db.Query("SELECT id, name, email, country, hour, price FROM tickets;")
	if err != nil {
		err = apperrors.ErrQueryDB
		return
	}
	defer rows.Close()

	t = make(map[int]internal.Ticket)
	for rows.Next() {
		var ticket internal.Ticket
		var attrs internal.TicketAttributes
		err = rows.Scan(&ticket.Id, &attrs.Name, &attrs.Email, &attrs.Country, &attrs.Hour, &attrs.Price)
		if err != nil {
			err = apperrors.ErrScanDB
			return
		}
		ticket.Attributes = attrs
		t[ticket.Id] = ticket
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
