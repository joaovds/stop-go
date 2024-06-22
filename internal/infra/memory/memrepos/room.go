package memrepos

import (
	"context"

	"github.com/joaovds/stop-go/internal/domain/player"
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
				Player: &player.Player{
					ID:        playerValue.ID,
					Nickname:  playerValue.Nickname,
					CreatedAt: playerValue.CreatedAt,
					UpdatedAt: playerValue.UpdatedAt,
				},
				Score:    playerValue.Score,
				Role:     room.PlayerRole(playerValue.Role),
				JoinedAt: playerValue.JoinedAt,
			}
		}

		result = append(result, &room.Room{
			ID:           roomValue.ID,
			Name:         roomValue.Name,
			Code:         roomValue.Code,
			MaxPlayers:   roomValue.Max,
			MinPlayers:   roomValue.Min,
			TotalPlayers: len(roomValue.Players),
			HostID:       roomValue.HostID,
			Players:      players,
			CreatedAt:    roomValue.CreatedAt,
			UpdatedAt:    roomValue.UpdatedAt,
		})
	}

	return result, nil
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
