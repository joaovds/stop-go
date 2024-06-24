package roomuc

import (
	"context"
	"time"

	"github.com/joaovds/stop-go/internal/domain/player"
	"github.com/joaovds/stop-go/internal/domain/room"
	"github.com/joaovds/stop-go/pkg/errs"
)

type FindByID struct {
	repo room.Repository
}

func NewFindByID(repo room.Repository) *FindByID {
	return &FindByID{repo: repo}
}

// ----- ... -----

func (f *FindByID) Execute(ctx context.Context, id string) (*findByIDOutput, *errs.Error) {
	rooms, err := f.repo.FindByID(ctx, id)
	if err.IsError() {
		return nil, err
	}

	return roomToOutput(rooms), nil
}

// ----- ... -----

type findByIDOutput struct {
	ID           string           `json:"id"`
	Name         string           `json:"name"`
	Code         string           `json:"code"`
	MaxPlayers   int              `json:"max_players"`
	MinPlayers   int              `json:"min_players"`
	Players      []*playersOutput `json:"players"`
	TotalPlayers int              `json:"total_players"`
	IsFull       bool             `json:"is_full"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
}

type playersOutput struct {
	ID        string    `json:"id"`
	Nickname  string    `json:"nickname"`
	Role      string    `json:"role"`
	Score     int       `json:"score"`
	JoinedAt  time.Time `json:"joined_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func roomToOutput(room *room.Room) *findByIDOutput {
	players := make([]*playersOutput, 0)
	for _, player := range room.Players {
		players = append(players, playerToOutput(player))
	}

	return &findByIDOutput{
		ID:           room.ID,
		Name:         room.Name,
		Code:         room.Code,
		MaxPlayers:   room.MaxPlayers,
		MinPlayers:   room.MinPlayers,
		Players:      players,
		TotalPlayers: room.TotalPlayers,
		IsFull:       room.IsFull(),
		CreatedAt:    room.CreatedAt,
		UpdatedAt:    room.UpdatedAt,
	}
}

func playerToOutput(player *player.Player) *playersOutput {
	return &playersOutput{
		ID:        player.ID,
		Nickname:  player.Nickname,
		Role:      player.Role.String(),
		Score:     player.Score,
		JoinedAt:  player.JoinedAt,
		CreatedAt: player.CreatedAt,
		UpdatedAt: player.UpdatedAt,
	}
}
