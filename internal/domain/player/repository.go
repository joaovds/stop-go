package player

import (
	"context"

	"github.com/joaovds/stop-go/pkg/errs"
)

type Repository interface {
	FindAll(ctx context.Context) ([]*Player, *errs.Error)
	NicknameExists(ctx context.Context, nickname string) (bool, *errs.Error)
	Exists(ctx context.Context, id string) (bool, *errs.Error)
	Create(ctx context.Context, player *Player) *errs.Error
}
