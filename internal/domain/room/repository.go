package room

import (
	"context"

	"github.com/joaovds/stop-go/pkg/errs"
)

type Repository interface {
	FindAll(ctx context.Context) ([]*Room, *errs.Error)
}
