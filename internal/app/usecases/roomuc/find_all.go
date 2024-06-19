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
	return new(FindAllOutput), nil
}

// ----- ... -----

type FindAllOutput []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
