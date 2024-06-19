package routes

import (
	"net/http"

	"github.com/joaovds/stop-go/internal/presentation/rest/handlers"
)

type RoomRoutes struct {
	muxV1    *http.ServeMux
	handlers *handlers.RoomHandlers
}

func NewRoomRoutes(muxV1 *http.ServeMux) *RoomRoutes {
	return &RoomRoutes{
		muxV1:    muxV1,
		handlers: handlers.NewRoomHandlers(),
	}
}

// ----- ... -----

func (r *RoomRoutes) RegisterRoutes() {
	r.muxV1.HandleFunc("/room", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello from room route!"))
	})
}
