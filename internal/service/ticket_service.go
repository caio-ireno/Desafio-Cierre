package service

import (
	"app/internal/domain"
	"context"
)

func NewServiceTicketDefault(rp domain.RepositoryTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

type ServiceTicketDefault struct {
	rp domain.RepositoryTicket
}

func (s *ServiceTicketDefault) GetAll(ctx context.Context) (ticket map[int]domain.Ticket, err error) {
	ticket, err = s.rp.GetAll(ctx)
	if err != nil {
		return
	}
	return
}

func (s *ServiceTicketDefault) AddCsv(ctx context.Context, csv map[int]domain.Ticket) (total int, err error) {
	total, err = s.rp.AddCsv(ctx, csv)
	if err != nil {
		return
	}
	return
}

func (s *ServiceTicketDefault) GetById(ctx context.Context, id int) (ticket domain.Ticket, err error) {
	ticket, err = s.rp.GetById(ctx, id)
	if err != nil {
		return
	}
	return
}

func (s *ServiceTicketDefault) GetTotalAmountTickets(ctx context.Context) (total int, err error) {
	total, err = s.rp.GetTotalAmountTickets(ctx)
	if err != nil {
		return 0, err
	}
	return
}

func (s *ServiceTicketDefault) Update(ctx context.Context, ticket domain.TicketAttributesPatch, id int) (ticketUpdate domain.Ticket, err error) {
	ticketUpdate, err = s.rp.Update(ctx, ticket, id)
	if err != nil {
		return
	}
	return
}

func (s *ServiceTicketDefault) Create(ctx context.Context, ticket domain.TicketAttributes) (ticketCreated domain.Ticket, err error) {
	ticketCreated, err = s.rp.Create(ctx, ticket)
	if err != nil {
		return
	}
	return
}
