package memrepos

import (
	"context"

	"github.com/joaovds/stop-go/internal/domain/room"
	"github.com/joaovds/stop-go/internal/infra/memory/memdata"
	"github.com/joaovds/stop-go/pkg/errs"
)

type RoomRepository struct{}

func NewRoomRepository() *RoomRepository {
	return &RoomRepository{}
}

// ----- ... -----

func (r *RoomRepository) FindAll(ctx context.Context) ([]*room.Room, *errs.Error) {
	rooms := memdata.Rooms

	result := make([]*room.Room, 0)
	for _, roomValue := range rooms {
		players := make(map[string]*room.Player)
		for _, playerValue := range roomValue.Players {
			players[playerValue.ID] = &room.Player{
				Player:   playerValue.ToDomain(),
				Score:    playerValue.Score,
				Role:     room.PlayerRole(playerValue.Role),
				JoinedAt: playerValue.JoinedAt,
			}
		}

		result = append(result, roomValue.ToDomain())
	}

	return result, nil
}

// ----- ... -----

func (r *RoomRepository) FindByID(ctx context.Context, id string) (*room.Room, *errs.Error) {
	roomValue, exists := memdata.Rooms[id]
	if !exists {
		return nil, errs.NewError("room not found").SetStatus(404)
	}
	players := make(map[string]*room.Player)
	for _, playerValue := range roomValue.Players {
		players[playerValue.ID] = &room.Player{
			Player:   playerValue.ToDomain(),
			Score:    playerValue.Score,
			Role:     room.PlayerRole(playerValue.Role),
			JoinedAt: playerValue.JoinedAt,
		}
	}

	return roomValue.ToDomain(), nil
}

// ----- ... -----

func (r *RoomRepository) NameExists(ctx context.Context, name string) (bool, *errs.Error) {
	for _, roomValue := range memdata.Rooms {
		if roomValue.Name == name {
			return true, nil
		}
	}

	return false, nil
}

// ----- ... -----

func (r *RoomRepository) Create(ctx context.Context, roomParams *room.Room) *errs.Error {
	newValue := &memdata.Room{
		ID:        roomParams.ID,
		Name:      roomParams.Name,
		Code:      roomParams.Code,
		Max:       roomParams.MaxPlayers,
		Min:       roomParams.MinPlayers,
		HostID:    roomParams.HostID,
		Password:  "",
		IsPrivate: false,
		Players:   make(map[string]*memdata.Player, 0),
		CreatedAt: roomParams.CreatedAt,
		UpdatedAt: roomParams.UpdatedAt,
	}

	memdata.Rooms[roomParams.ID] = newValue
	return nil
}
