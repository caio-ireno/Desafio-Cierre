package repository

import (
	"app/internal/domain"
	"app/pkg/apperrors"
	"context"
	"database/sql"
	"fmt"
)

func NewRepositoryTicket(db *sql.DB) domain.RepositoryTicket {
	return &mysqlRepository{
		db: db,
	}
}

type mysqlRepository struct {
	db *sql.DB
}

func (r mysqlRepository) GetAll(ctx context.Context) (t map[int]domain.Ticket, err error) {
	rows, err := r.db.Query("SELECT id, name, email, country, hour, price FROM tickets;")
	if err != nil {
		err = apperrors.ErrQueryDB
		return
	}
	defer rows.Close()

	t = make(map[int]domain.Ticket)
	for rows.Next() {
		var ticket domain.Ticket
		var attrs domain.TicketAttributes
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

func (r mysqlRepository) GetById(ctx context.Context, id int) (t domain.Ticket, err error) {
	row := r.db.QueryRow("SELECT id, name, email, country, hour, price FROM tickets WHERE id = ?;", id)
	var attrs domain.TicketAttributes
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
	err = r.db.QueryRow("SELECT COUNT(id) FROM tickets;").Scan(&total)
	if total == 0 {
		err = apperrors.ErrEmptyData
		return
	}

	if err != nil {
		err = apperrors.ErrQueryDB
		return
	}
	return
}

func (r *mysqlRepository) AddCsv(ctx context.Context, csv map[int]domain.Ticket) (total int, err error) {
	for _, record := range csv {
		total++
		_, err = r.db.ExecContext(ctx,
			`INSERT INTO tickets (name, email, country, hour, price) VALUES (?, ?, ?, ?, ?)`,
			record.Attributes.Name,
			record.Attributes.Email,
			record.Attributes.Country,
			record.Attributes.Hour,
			record.Attributes.Price,
		)
		if err != nil {
			err = apperrors.ErrQueryDB
			return
		}
	}
	return
}

func (r *mysqlRepository) Update(ctx context.Context, ticket domain.TicketAttributesPatch, id int) (ticketUpdate domain.Ticket, err error) {

	return
}

func (r *mysqlRepository) Create(ctx context.Context, ticket domain.TicketAttributes) (ticketCreated domain.Ticket, err error) {

	return
}
