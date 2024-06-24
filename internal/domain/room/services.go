package room

import (
	"context"
	"time"

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

// ----- ... -----

func (s *Service) CreateRoom(ctx context.Context, creator *player.Player, roomParam *Room) *errs.Error {
	if err := roomParam.Validate(); err.IsError() {
		return err
	}
	if err := creator.Validate(); err.IsError() {
		return err
	}

	_, err := s.playerRepo.FindByID(ctx, creator.ID)
	if err.IsNil() {
		return player.ErrPlayerAlreadyExists
	}
	if exists, err := s.roomRepo.NameExists(ctx, roomParam.Name); err.IsError() {
		return err
	} else if exists {
		return ErrNameAlreadyExists
	}

	creator.Role = player.Host

	err = s.playerRepo.Create(ctx, creator)
	if err.IsError() {
		return err
	}
	err = s.roomRepo.Create(ctx, roomParam)
	if err.IsError() {
		return err
	} else {
		creator.JoinedAt = time.Now()
	}

	return nil
}
