package room

import (
	"context"

	"github.com/joaovds/stop-go/pkg/errs"
)

type Repository interface {
	FindAll(ctx context.Context) ([]*Room, *errs.Error)
	NameExists(ctx context.Context, name string) (bool, *errs.Error)
	Create(ctx context.Context, room *Room) *errs.Error
}
