package memrepos

import (
	"context"

	"github.com/joaovds/stop-go/internal/domain/room"
	"github.com/joaovds/stop-go/pkg/errs"
)

type RoomRepository struct{}

func NewRoomRepository() *RoomRepository {
	return &RoomRepository{}
}

// ----- ... -----

func (r *RoomRepository) FindAll(ctx context.Context) ([]*room.Room, *errs.Error) {
	return []*room.Room{
		{
			ID:   1,
			Name: "Room 1",
		},
	}, nil
}
