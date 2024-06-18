package memrepos

import (
	"sync"

	"github.com/joaovds/stop-go/internal/domain/entities"
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

func (r *RoomRepository) List() ([]*entities.Room, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	rooms := make([]*entities.Room, 0, len(r.rooms))
	for _, room := range r.rooms {
		rooms = append(rooms, room)
	}

	return rooms, nil
}
