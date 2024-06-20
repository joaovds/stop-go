package routes

import (
	"net/http"

	"github.com/joaovds/stop-go/internal/app/usecases/playeruc"
	"github.com/joaovds/stop-go/internal/infra/memory/memrepos"
	"github.com/joaovds/stop-go/internal/infra/sqlite"
	"github.com/joaovds/stop-go/internal/presentation/rest/handlers"
)

type PlayerRoutes struct {
	muxV1    *http.ServeMux
	handlers *handlers.PlayerHandlers
}

// ----- ... -----

func NewPlayerRoutes(muxV1 *http.ServeMux, db *sqlite.DB) *PlayerRoutes {
	playerRepo := memrepos.NewPlayerRepository(db)

	return &PlayerRoutes{
		muxV1: muxV1,
		handlers: handlers.NewPlayerHandlers(
			playeruc.NewFindAll(playerRepo),
			playeruc.NewCreate(playerRepo),
		),
	}
}

// ----- ... -----

func (p *PlayerRoutes) RegisterRoutes() {
	p.muxV1.HandleFunc("GET /players", p.handlers.FindAll)
	p.muxV1.HandleFunc("POST /players", p.handlers.Create)
}
