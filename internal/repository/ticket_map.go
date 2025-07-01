package repository

import (
	"app/internal"
	"app/pkg/apperrors"
	"context"
	"reflect"
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

func (r *RepositoryTicketMap) Update(ctx context.Context, ticket internal.TicketAttributesPatch, id int) (ticketUpdate internal.Ticket, err error) {

	v, ok := r.db[id]

	if !ok {
		err = apperrors.ErrNotFound
		return
	}
	// O pacote reflect do Go permite inspecionar e manipular valores em tempo de execução, mesmo sem saber seus tipos exatos em tempo de compilação.
	// No seu código, ele é usado para atualizar dinamicamente apenas os campos enviados no PATCH, sem precisar de vários
	orig := &v.Attributes
	patchVal := reflect.ValueOf(ticket)     // struct PATCH recebido (com ponteiros)
	origVal := reflect.ValueOf(orig).Elem() // struct original a ser atualizado

	for i := 0; i < patchVal.NumField(); i++ {
		patchField := patchVal.Field(i)
		if !patchField.IsNil() {
			origField := origVal.Field(i)
			origField.Set(reflect.Indirect(patchField))
		}
	}

	r.db[id] = v
	ticketUpdate = v
	return
}

func (r *RepositoryTicketMap) Create(ctx context.Context, ticket internal.TicketAttributes) (ticketCreated internal.Ticket, err error) {
	lastId := len(r.db)
	NewId := lastId + 1

	ticketCreated = internal.Ticket{
		Id:         NewId,
		Attributes: ticket,
	}

	r.db[ticketCreated.Id] = ticketCreated
	return
}
