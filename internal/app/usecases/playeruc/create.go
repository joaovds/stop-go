package playeruc

import (
	"context"

	"github.com/joaovds/stop-go/internal/domain/player"
	"github.com/joaovds/stop-go/pkg/errs"
)

type Create struct {
	repo player.Repository
}

func NewCreate(repo player.Repository) *Create {
	return &Create{repo: repo}
}

// ----- ... -----

func (c *Create) Execute(ctx context.Context, params *CreateInput) *errs.Error {
	input, err := inputToPlayer(params)
	if err.IsError() {
		return err
	}

	err = c.repo.Create(ctx, input)
	if err.IsError() {
		return err
	}

	return nil
}

// ----- ... -----

type CreateInput struct {
	Nickname string `json:"nickname"`
}

func inputToPlayer(input *CreateInput) (*player.Player, *errs.Error) {
	res, err := player.NewPlayer(input.Nickname)
	if err.IsError() {
		return nil, err
	}

	return res, nil
}
