package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/joaovds/stop-go/internal/app/usecases/playeruc"
	"github.com/joaovds/stop-go/internal/presentation/rest/helpers"
	"github.com/joaovds/stop-go/internal/presentation/rest/response"
)

type PlayerHandlers struct {
	findAllUseCase *playeruc.FindAll
	createUseCase  *playeruc.Create
}

func NewPlayerHandlers(
	findAllUseCase *playeruc.FindAll,
	createUseCase *playeruc.Create,
) *PlayerHandlers {
	return &PlayerHandlers{
		findAllUseCase: findAllUseCase,
		createUseCase:  createUseCase,
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

func (h *PlayerHandlers) Create(w http.ResponseWriter, req *http.Request) {
	var player playeruc.CreateInput
	if err := json.NewDecoder(req.Body).Decode(&player); err != nil {
		helpers.NewBadRequestError().WriteJson(w)
		return
	}

	err := h.createUseCase.Execute(req.Context(), &player)
	if err.IsError() {
		response.NewHttpResError(err.Status, err.Error()).WriteJson(w)
		return
	}

	helpers.NewCreated(nil).WriteJson(w)
}
