package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/joaovds/stop-go/internal/app/usecases/playeruc"
)

type PlayerHandlers struct {
	findAllUseCase *playeruc.FindAll
}

func NewPlayerHandlers(
	findAllUseCase *playeruc.FindAll,
) *PlayerHandlers {
	return &PlayerHandlers{
		findAllUseCase: findAllUseCase,
	}
}

// ----- ... -----

func (h *PlayerHandlers) FindAll(w http.ResponseWriter, req *http.Request) {
	players, err := h.findAllUseCase.Execute(req.Context())
	if err.IsError() {
		w.WriteHeader(err.Status)
		json.NewEncoder(w).Encode(err.Message)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(players)
}
