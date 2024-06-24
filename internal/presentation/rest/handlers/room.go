package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/joaovds/stop-go/internal/app/usecases/roomuc"
	"github.com/joaovds/stop-go/internal/presentation/rest/helpers"
	"github.com/joaovds/stop-go/internal/presentation/rest/response"
)

type RoomHandlers struct {
	findAllUseCase  *roomuc.FindAll
	findByIDUseCase *roomuc.FindByID
	createUseCase   *roomuc.Create
}

func NewRoomHandlers(
	findAllUseCase *roomuc.FindAll,
	findByIDUseCase *roomuc.FindByID,
	createUseCase *roomuc.Create,
) *RoomHandlers {
	return &RoomHandlers{
		findAllUseCase:  findAllUseCase,
		findByIDUseCase: findByIDUseCase,
		createUseCase:   createUseCase,
	}
}

// ----- ... -----

func (h *RoomHandlers) FindAll(w http.ResponseWriter, req *http.Request) {
	rooms, err := h.findAllUseCase.Execute(req.Context())
	if err.IsError() {
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(err.Message)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rooms)
}

// ----- ... -----

func (h *RoomHandlers) FindByID(w http.ResponseWriter, req *http.Request) {
	roomID := req.PathValue("id")

	roomData, err := h.findByIDUseCase.Execute(req.Context(), roomID)
	if err.IsError() {
		response.NewHttpResError(err.Status, err.Error()).WriteJson(w)
		return
	}

	helpers.NewOk(roomData).WriteJson(w)
}

// ----- ... -----

func (h *RoomHandlers) Create(w http.ResponseWriter, req *http.Request) {
	var room roomuc.CreateInput
	if err := json.NewDecoder(req.Body).Decode(&room); err != nil {
		helpers.NewBadRequestError().WriteJson(w)
		return
	}

	err := h.createUseCase.Execute(req.Context(), &room)
	if err.IsError() {
		response.NewHttpResError(err.Status, err.Error()).WriteJson(w)
		return
	}

	helpers.NewCreated(nil).WriteJson(w)
}

// ----- ... -----
