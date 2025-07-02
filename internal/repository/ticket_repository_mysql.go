package repository

import (
	"app/internal"
	"app/pkg/apperrors"
	"context"
	"database/sql"
	"fmt"
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

func (r mysqlRepository) GetById(ctx context.Context, id int) (t internal.Ticket, err error) {
	row := r.db.QueryRow("SELECT id, name, email, country, hour, price FROM tickets WHERE id = ?;", id)
	var attrs internal.TicketAttributes
	err = row.Scan(&t.Id, &attrs.Name, &attrs.Email, &attrs.Country, &attrs.Hour, &attrs.Price)
	if err == sql.ErrNoRows {
		fmt.Println("Erro No rows", err)
		err = apperrors.ErrNotFound
		return
	}
	if err != nil {
		fmt.Println("Erro scan", err)
		return t, apperrors.ErrScanDB
	}

	t.Attributes = attrs
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
