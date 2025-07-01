package repository

import (
	"app/internal"
	"app/pkg/apperrors"
	"context"
)

func NewRepositoryTicketMap(dbFile map[int]internal.Ticket, lastId int) *RepositoryTicketMap {
	defaultDb := make(map[int]internal.Ticket)

	if dbFile != nil {
		defaultDb = dbFile
	}

	return &RepositoryTicketMap{
		db:     defaultDb,
		lastId: lastId,
	}
}

type RepositoryTicketMap struct {
	db     map[int]internal.Ticket
	lastId int
}

func (r *RepositoryTicketMap) GetAll(ctx context.Context) (t map[int]internal.Ticket, err error) {
	t = make(map[int]internal.Ticket, len(r.db))
	for k, v := range r.db {
		t[k] = v
	}

	if len(t) == 0 {
		err = apperrors.ErrEmptyData
	}

	return
}

func (r *RepositoryTicketMap) GetTicketByDestinationCountry(ctx context.Context, country string) (t map[int]internal.TicketAttributes, err error) {
	t = make(map[int]internal.TicketAttributes)
	for k, v := range r.db {
		if v.Attributes.Country == country {
			t[k] = v.Attributes
		}
	}

	return
}

func (r *RepositoryTicketMap) GetTotalAmountTickets(ctx context.Context) (total int, err error) {
	for range r.db {
		total++
	}
	return
}
