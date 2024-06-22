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

// ----- ... -----

func (s *Service) CreateRoom(ctx context.Context, creatorID string, roomParam *Room) *errs.Error {
	playerRes, err := s.playerRepo.FindByID(ctx, creatorID)
	if err.IsError() {
		return err
	}

	if exists, err := s.roomRepo.NameExists(ctx, roomParam.Name); err.IsError() {
		return err
	} else if exists {
		return ErrNameAlreadyExists
	}

	if p, err := NewPlayer(playerRes, Host); err.IsError() {
		return err
	} else {
		roomParam.Players[playerRes.ID] = p
	}

	err = s.roomRepo.Create(ctx, roomParam)
	if err.IsError() {
		return err
	}

	return nil
}
