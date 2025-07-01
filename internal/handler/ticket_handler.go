package handler

import (
	"app/internal"
	"app/pkg/apperrors"
	"errors"
	"net/http"

	"github.com/bootcamp-go/web/response"
)

func NewHandlerTicketDefault(sv internal.ServiceTicket) *TicketDefault {

	return &TicketDefault{sv: sv}

}

type TicketDefault struct {
	sv internal.ServiceTicket
}

func (h *TicketDefault) GetTotalAmountTickets() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		total, err := h.sv.GetTotalAmountTickets(ctx)

		if err != nil {
			response.JSON(w, http.StatusBadRequest, "")
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"data": total,
		})

	}

}

func (h *TicketDefault) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		tickets, err := h.sv.GetAll(ctx)

		if err != nil {
			if errors.Is(err, apperrors.ErrEmptyData) {
				response.JSON(w, http.StatusBadRequest, err.Error())
			}
			response.JSON(w, http.StatusBadRequest, "Somethings wrong!!")
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"data": tickets,
		})

	}

}
