package roomuc

import (
	"context"

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

type FindAllOutput []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func roomsToOutput(rooms []*room.Room) *FindAllOutput {
	output := make(FindAllOutput, len(rooms))

	for i, r := range rooms {
		output[i] = struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		}{
			ID:   r.ID,
			Name: r.Name,
		}
	}

	return &output
}
