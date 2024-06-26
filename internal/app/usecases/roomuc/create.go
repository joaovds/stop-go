package roomuc

import (
	"context"

	"github.com/joaovds/stop-go/internal/domain/player"
	"github.com/joaovds/stop-go/internal/domain/room"
	"github.com/joaovds/stop-go/internal/infra/providers"
	"github.com/joaovds/stop-go/pkg/errs"
)

type Create struct {
	repo    room.Repository
	service *room.Service
}

func NewCreate(repo room.Repository, playerRepo player.Repository) *Create {
	service := room.NewService(repo, playerRepo)
	return &Create{repo: repo, service: service}
}

// ----- ... -----

func (c *Create) Execute(ctx context.Context, params *CreateInput) *errs.Error {
	if err := params.Validate(); err != nil {
		return err
	}

	roomInput, playerInput, err := inputToRoom(params)
	if err.IsError() {
		return err
	}

	err = c.service.CreateRoom(ctx, playerInput, roomInput)
	if err.IsError() {
		return err
	}

	return nil
}

// ----- ... -----

type CreateInput struct {
	Creator    Creator `json:"creator"`
	Name       string  `json:"name"`
	MaxPlayers int     `json:"max_players"`
	MinPlayers int     `json:"min_players"`
}

type Creator struct {
	Nickname string `json:"nickname"`
}

func (i *CreateInput) Validate() *errs.Error {
	if i.Creator.Nickname == "" {
		return errs.NewError("creator nickname is required").SetStatus(400)
	}
	if i.Name == "" {
		return errs.NewError("name is required").SetStatus(400)
	}
	if i.MaxPlayers == 0 {
		return errs.NewError("max_players is required").SetStatus(400)
	}
	if i.MinPlayers == 0 {
		return errs.NewError("min_players is required").SetStatus(400)
	}
	return nil
}

func inputToRoom(input *CreateInput) (*room.Room, *player.Player, *errs.Error) {
	p, err := player.NewPlayer(providers.NewID(), input.Creator.Nickname, player.Host)
	if err.IsError() {
		return nil, nil, err
	}
	if err := p.Validate(); err.IsError() {
		return nil, nil, err
	}

	res, err := room.NewRoom(
		providers.NewID(),
		input.Name,
		input.MaxPlayers,
		input.MinPlayers,
		p,
	)
	if err.IsError() {
		return nil, nil, err
	}
	if err := res.Validate(); err.IsError() {
		return nil, nil, err
	}

	return res, p, nil
}
