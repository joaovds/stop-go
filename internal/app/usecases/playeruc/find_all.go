package playeruc

import (
	"context"

	"github.com/joaovds/stop-go/internal/domain/player"
	"github.com/joaovds/stop-go/pkg/errs"
)

type FindAll struct {
	repo player.Repository
}

func NewFindAll(repo player.Repository) *FindAll {
	return &FindAll{repo: repo}
}

// ----- ... -----

func (f *FindAll) Execute(ctx context.Context) (*FindAllOutput, *errs.Error) {
	rooms, err := f.repo.FindAll(ctx)
	if err.IsError() {
		return nil, err
	}

	return playersToOutput(rooms), nil
}

// ----- ... -----

type findAllOutputItem struct {
	ID       string `json:"id"`
	Nickname string `json:"name"`
}

type FindAllOutput []findAllOutputItem

func playersToOutput(players []*player.Player) *FindAllOutput {
	output := make(FindAllOutput, len(players))

	for i, p := range players {
		output[i] = findAllOutputItem{
			ID:       p.ID,
			Nickname: p.Nickname,
		}
	}

	return &output
}
