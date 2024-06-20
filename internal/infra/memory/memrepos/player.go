package memrepos

import (
	"context"

	"github.com/joaovds/stop-go/internal/domain/player"
	"github.com/joaovds/stop-go/pkg/errs"
)

type PlayerRepository struct{}

func NewPlayerRepository() *PlayerRepository {
	return &PlayerRepository{}
}

// ----- ... -----

func (p *PlayerRepository) FindAll(ctx context.Context) ([]*player.Player, *errs.Error) {
	return []*player.Player{
		player.NewPlayer("Player 1"),
	}, nil
}
