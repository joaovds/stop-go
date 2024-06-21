package roomuc

import (
	"context"
	"time"

	"github.com/joaovds/stop-go/internal/domain/room"
	"github.com/joaovds/stop-go/pkg/errs"
)

type FindAll struct {
	repo room.Repository
}

func NewFindAll(repo room.Repository) *FindAll {
	return &FindAll{repo: repo}
}

// ----- ... -----

func (f *FindAll) Execute(ctx context.Context) (*FindAllOutput, *errs.Error) {
	rooms, err := f.repo.FindAll(ctx)
	if err.IsError() {
		return nil, err
	}

	return roomsToOutput(rooms), nil
}

// ----- ... -----

type findAllOutputItem struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Code         string    `json:"code"`
	MaxPlayers   int       `json:"max_players"`
	MinPlayers   int       `json:"min_players"`
	TotalPlayers int       `json:"total_players"`
	IsFull       bool      `json:"is_full"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type FindAllOutput []findAllOutputItem

func roomsToOutput(rooms []*room.Room) *FindAllOutput {
	output := make(FindAllOutput, len(rooms))

	for i, r := range rooms {
		output[i] = findAllOutputItem{
			ID:           r.ID,
			Name:         r.Name,
			Code:         r.Code,
			MaxPlayers:   r.MaxPlayers,
			MinPlayers:   r.MinPlayers,
			TotalPlayers: r.TotalPlayers,
			IsFull:       r.TotalPlayers == r.MaxPlayers,
			CreatedAt:    r.CreatedAt,
			UpdatedAt:    r.UpdatedAt,
		}
	}

	return &output
}
