package routes

import (
	"net/http"

	"github.com/joaovds/stop-go/internal/app/usecases/roomuc"
	"github.com/joaovds/stop-go/internal/infra/memory/memrepos"
	"github.com/joaovds/stop-go/internal/infra/sqlite"
	"github.com/joaovds/stop-go/internal/infra/sqlite/sqliterepos"
	"github.com/joaovds/stop-go/internal/presentation/rest/handlers"
)

type RoomRoutes struct {
	muxV1    *http.ServeMux
	handlers *handlers.RoomHandlers
}

// ----- ... -----

func NewRoomRoutes(muxV1 *http.ServeMux, db *sqlite.DB) *RoomRoutes {
	roomRepo := memrepos.NewRoomRepository()
	playerRepo := sqliterepos.NewPlayerRepository(db)

	return &RoomRoutes{
		muxV1: muxV1,
		handlers: handlers.NewRoomHandlers(
			roomuc.NewFindAll(roomRepo),
			roomuc.NewCreate(roomRepo, playerRepo),
		),
	}
}

// ----- ... -----

func (r *RoomRoutes) RegisterRoutes() {
	r.muxV1.HandleFunc("GET /rooms", r.handlers.FindAll)
	r.muxV1.HandleFunc("POST /rooms", r.handlers.Create)
}
