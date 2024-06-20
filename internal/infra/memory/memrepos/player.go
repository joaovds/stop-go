package memrepos

import (
	"context"

	"github.com/joaovds/stop-go/internal/domain/player"
	"github.com/joaovds/stop-go/pkg/errs"
)

type PlayerRepository struct {
	players []*player.Player
}

func NewPlayerRepository() *PlayerRepository {
	return &PlayerRepository{}
}

// ----- ... -----

func (p *PlayerRepository) FindAll(ctx context.Context) ([]*player.Player, *errs.Error) {
	return p.players, nil
}

// ----- ... -----

func (p *PlayerRepository) Create(ctx context.Context, player *player.Player) *errs.Error {
	p.players = append(p.players, player)
	return nil
}
