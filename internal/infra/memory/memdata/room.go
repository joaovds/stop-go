package memdata

import (
	"time"

	"github.com/joaovds/stop-go/internal/domain/player"
	"github.com/joaovds/stop-go/internal/domain/room"
)

var Rooms = make(map[string]*Room)

type Room struct {
	ID        string
	Name      string
	Code      string
	Max       int
	Min       int
	HostID    string
	Password  string
	IsPrivate bool
	Players   map[string]*Player
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ----- ... -----

func (r *Room) ToDomain() *room.Room {
	players := make(map[string]*room.Player)
	for _, playerValue := range r.Players {
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

	result := &room.Room{
		ID:           r.ID,
		Name:         r.Name,
		Code:         r.Code,
		MaxPlayers:   r.Max,
		MinPlayers:   r.Min,
		TotalPlayers: len(r.Players),
		HostID:       r.HostID,
		Players:      players,
		CreatedAt:    r.CreatedAt,
		UpdatedAt:    r.UpdatedAt,
	}

	return result
}
