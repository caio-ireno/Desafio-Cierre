package service

import (
	"app/internal"
	"context"
)

func NewServiceTicketDefault(rp internal.RepositoryTicket) *ServiceTicketDefault {
	return &ServiceTicketDefault{
		rp: rp,
	}
}

type ServiceTicketDefault struct {
	rp internal.RepositoryTicket
}

func (s *ServiceTicketDefault) GetAll(ctx context.Context) (ticket map[int]internal.Ticket, err error) {
	ticket, err = s.rp.GetAll(ctx)
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

func (s *ServiceTicketDefault) Update(ctx context.Context, ticket internal.TicketAttributesPatch, id int) (ticketUpdate internal.Ticket, err error) {
	ticketUpdate, err = s.rp.Update(ctx, ticket, id)
	if err != nil {
		return
	}
	return
}
