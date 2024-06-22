package routes

import (
	"net/http"

	"github.com/joaovds/stop-go/internal/presentation/rest/handlers"
)

type PlayerRoutes struct {
	muxV1    *http.ServeMux
	handlers *handlers.PlayerHandlers
}

// ----- ... -----

func NewPlayerRoutes(muxV1 *http.ServeMux) *PlayerRoutes {
	return &PlayerRoutes{
		muxV1:    muxV1,
		handlers: handlers.NewPlayerHandlers(),
	}
}

// ----- ... -----

func (p *PlayerRoutes) RegisterRoutes() {}
