package routes

import (
	"net/http"

	"github.com/joaovds/stop-go/internal/app/usecases/playeruc"
	"github.com/joaovds/stop-go/internal/infra/memory/memrepos"
	"github.com/joaovds/stop-go/internal/presentation/rest/handlers"
)

type PlayerRoutes struct {
	muxV1    *http.ServeMux
	handlers *handlers.PlayerHandlers
}

// ----- ... -----

func NewPlayerRoutes(muxV1 *http.ServeMux) *PlayerRoutes {
	playerRepo := memrepos.NewPlayerRepository()

	findAllUseCase := playeruc.NewFindAll(playerRepo)

	return &PlayerRoutes{
		muxV1: muxV1,
		handlers: handlers.NewPlayerHandlers(
			findAllUseCase,
		),
	}
}

// ----- ... -----

func (p *PlayerRoutes) RegisterRoutes() {
	p.muxV1.HandleFunc("/players", p.handlers.FindAll)
}
