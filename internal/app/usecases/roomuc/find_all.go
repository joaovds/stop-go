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

type findAllOutputItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type FindAllOutput []findAllOutputItem

func roomsToOutput(rooms []*room.Room) *FindAllOutput {
	output := make(FindAllOutput, len(rooms))

	for i, r := range rooms {
		output[i] = findAllOutputItem{
			ID:   r.ID,
			Name: r.Name,
		}
	}

	return &output
}
