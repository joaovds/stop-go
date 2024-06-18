package memrepos

import (
	"sync"

	"github.com/joaovds/stop-go/internal/domain/entities"
	"github.com/joaovds/stop-go/internal/shared/errs"
)

type RoomRepository struct {
	rooms map[string]*entities.Room
	mutex *sync.Mutex
}

func NewRoomRepository() *RoomRepository {
	return &RoomRepository{
		rooms: make(map[string]*entities.Room),
		mutex: &sync.Mutex{},
	}
}

// ----- ... -----

func (r *RoomRepository) List() ([]*entities.Room, *errs.Error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	rooms := make([]*entities.Room, 0, len(r.rooms))
	for _, room := range r.rooms {
		rooms = append(rooms, room)
	}

	rooms = append(rooms, &entities.Room{
		ID:   1,
		Name: "Room 1",
		Code: "ROOM_1",
	})

	return rooms, nil
}
