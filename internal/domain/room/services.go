package room

import (
	"context"

	"github.com/joaovds/stop-go/internal/domain/player"
	"github.com/joaovds/stop-go/pkg/errs"
)

type Service struct {
	roomRepo   Repository
	playerRepo player.Repository
}

func NewService(roomRepo Repository, playerRepo player.Repository) *Service {
	return &Service{roomRepo: roomRepo, playerRepo: playerRepo}
}

// ----- ... -----

func (s *Service) PlayerExists(ctx context.Context, playerID string) (bool, *errs.Error) {
	return s.playerRepo.Exists(ctx, playerID)
}
