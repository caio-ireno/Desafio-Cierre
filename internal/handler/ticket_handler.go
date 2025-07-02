package handler

import (
	"app/internal/domain"
	"app/internal/loader"
	"app/pkg/apperrors"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

func NewHandlerTicketDefault(sv domain.ServiceTicket) *TicketDefault {
	return &TicketDefault{sv: sv}
}

type TicketDefault struct {
	sv domain.ServiceTicket
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
func (h *TicketDefault) GetById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		ticket, err := h.sv.GetById(ctx, id)

		if err != nil {
			if errors.Is(err, apperrors.ErrNotFound) {
				response.JSON(w, http.StatusNotFound, err.Error())
				return
			}
			response.JSON(w, http.StatusBadRequest, err.Error())
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"data": ticket,
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

func (h *TicketDefault) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		idStr := chi.URLParam(r, "id")
		id, err := strconv.Atoi(idStr)

		if err != nil {
			// ID inválido
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		var reqBody domain.TicketAttributesPatch

		err = json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Invalid data")
			return
		}

		tickets, err := h.sv.Update(ctx, reqBody, id)
		if err != nil {
			if errors.Is(err, apperrors.ErrEmptyData) {
				response.JSON(w, http.StatusBadRequest, err.Error())
				return
			}
			response.JSON(w, http.StatusBadRequest, "Somethings wrong!!")
			return
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"data": tickets,
		})

	}

}

func (h *TicketDefault) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()

		var reqBody domain.TicketAttributes

		err := json.NewDecoder(r.Body).Decode(&reqBody)

		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Invalid data")
			return
		}

		err = reqBody.Validate()

		if err != nil {
			response.JSON(w, http.StatusBadRequest, err.Error())
			return
		}

		tickets, err := h.sv.Create(ctx, reqBody)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Somethings wrong!!")
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"data": tickets,
		})

	}

}

func (h *TicketDefault) AddCsv() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		ld := loader.NewLoaderTicketCSV("docs/db/tickets.csv")

		tickets, err := ld.Load()
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Error to load CSV!!")
			return
		}
		total, err := h.sv.AddCsv(ctx, tickets)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, "Somethings wrong!!")
			return
		}

		response.JSON(w, http.StatusOK, map[string]any{
			"message": "Tickets imported successfully",
			"total":   total,
		})

	}

}
