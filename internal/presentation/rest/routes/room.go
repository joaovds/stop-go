package routes

import (
	"net/http"

	"github.com/joaovds/stop-go/internal/app/usecases/roomuc"
	"github.com/joaovds/stop-go/internal/infra/memory/memrepos"
	"github.com/joaovds/stop-go/internal/presentation/rest/handlers"
)

type RoomRoutes struct {
	muxV1    *http.ServeMux
	handlers *handlers.RoomHandlers
}

func NewRoomRoutes(muxV1 *http.ServeMux) *RoomRoutes {
	roomRepo := memrepos.NewRoomRepository()

	findAllUseCase := roomuc.NewFindAll(roomRepo)

	return &RoomRoutes{
		muxV1: muxV1,
		handlers: handlers.NewRoomHandlers(
			findAllUseCase,
		),
	}
}

// ----- ... -----

func (r *RoomRoutes) RegisterRoutes() {
	r.muxV1.HandleFunc("/room", r.handlers.FindAll)
}
