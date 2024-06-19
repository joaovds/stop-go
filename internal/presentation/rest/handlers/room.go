package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/joaovds/stop-go/internal/app/usecases/roomuc"
)

type RoomHandlers struct {
	findAllUseCase *roomuc.FindAll
}

func NewRoomHandlers(
	findAllUseCase *roomuc.FindAll,
) *RoomHandlers {
	return &RoomHandlers{
		findAllUseCase: findAllUseCase,
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
