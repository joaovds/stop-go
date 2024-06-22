package memrepos

import (
	"context"

	"github.com/joaovds/stop-go/internal/domain/player"
	"github.com/joaovds/stop-go/internal/infra/memory/memdata"
	"github.com/joaovds/stop-go/pkg/errs"
)

type PlayerRepository struct{}

func NewPlayerRepository() *PlayerRepository {
	return &PlayerRepository{}
}

// ----- ... -----

func (p *PlayerRepository) FindByID(ctx context.Context, id string) (*player.Player, *errs.Error) {
	player, exists := memdata.Players[id]
	if !exists {
		return nil, errs.NewError("player not found").SetStatus(404)
	}

	return player.ToDomain(), nil
}

// ----- ... -----

func (p *PlayerRepository) NicknameExists(ctx context.Context, nickname string) (bool, *errs.Error) {
	for _, p := range memdata.Players {
		if p.Nickname == nickname {
			return true, nil
		}
	}

	return false, nil
}

// ----- ... -----

func (p *PlayerRepository) Exists(ctx context.Context, id string) (bool, *errs.Error) {
	_, exists := memdata.Players[id]
	return exists, nil
}
