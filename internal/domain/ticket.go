package domain

import (
	"context"
	"errors"
)

// TicketAttributes is an struct that represents a ticket
type TicketAttributes struct {
	// Name represents the name of the owner of the ticket
	Name string `json:"name"`
	// Email represents the email of the owner of the ticket
	Email string `json:"email"`
	// Country represents the destination country of the ticket
	Country string `json:"country"`
	// Hour represents the hour of the ticket
	Hour string `json:"hour"`
	// Price represents the price of the ticket
	Price float64 `json:"price"`
}

// Ticket represents a ticket
type Ticket struct {
	// Id represents the id of the ticket
	Id int `json:"id"`
	// Attributes represents the attributes of the ticket
	Attributes TicketAttributes `json:"attributes"`
}

type TicketAttributesPatch struct {
	Name    *string  `json:"name,omitempty"`
	Email   *string  `json:"email,omitempty"`
	Country *string  `json:"country,omitempty"`
	Hour    *string  `json:"hour,omitempty"`
	Price   *float64 `json:"price,omitempty"`
}

func (ta *TicketAttributes) Validate() error {
	if ta.Name == "" {
		return errors.New("name is required")
	}
	if ta.Email == "" {
		return errors.New("email is required")
	}
	if ta.Country == "" {
		return errors.New("country is required")
	}
	if ta.Hour == "" {
		return errors.New("hour is required")
	}
	if ta.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	return nil
}

type RepositoryTicket interface {
	GetById(ctx context.Context, id int) (t Ticket, err error)
	GetAll(ctx context.Context) (t map[int]Ticket, err error)
	GetTotalAmountTickets(ctx context.Context) (total int, err error)

	AddCsv(ctx context.Context, csv map[int]Ticket) (total int, err error)

	Update(ctx context.Context, ticket TicketAttributesPatch, id int) (t Ticket, err error)
	Create(ctx context.Context, ticket TicketAttributes) (t Ticket, err error)
}

type ServiceTicket interface {
	GetById(ctx context.Context, id int) (t Ticket, err error)
	GetAll(ctx context.Context) (t map[int]Ticket, err error)

	AddCsv(ctx context.Context, csv map[int]Ticket) (total int, err error)

	GetTotalAmountTickets(ctx context.Context) (total int, err error)

	Update(ctx context.Context, ticket TicketAttributesPatch, id int) (t Ticket, err error)
	Create(ctx context.Context, ticket TicketAttributes) (t Ticket, err error)
}
