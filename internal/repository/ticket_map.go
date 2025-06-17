package repository

import (
	"app/internal"
	"context"
)

func NewRepositoryTicketMap(dbFile map[int]internal.TicketAttributes, lastId int) *RepositoryTicketMap {
	return &RepositoryTicketMap{
		db:     dbFile,
		lastId: lastId,
	}
}

type RepositoryTicketMap struct {
	db map[int]internal.TicketAttributes

	lastId int
}

func (r *RepositoryTicketMap) Get(ctx context.Context) (t map[int]internal.TicketAttributes, err error) {
	t = make(map[int]internal.TicketAttributes, len(r.db))
	for k, v := range r.db {
		t[k] = v
	}

	return
}

func (r *RepositoryTicketMap) GetTicketsByDestinationCountry(ctx context.Context, country string) (t map[int]internal.TicketAttributes, err error) {
	t = make(map[int]internal.TicketAttributes)
	for k, v := range r.db {
		if v.Country == country {
			t[k] = v
		}
	}

	return
}
