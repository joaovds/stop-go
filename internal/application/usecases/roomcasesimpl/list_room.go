package roomcasesimpl

import (
	"github.com/joaovds/stop-go/internal/domain/repositories"
	"github.com/joaovds/stop-go/internal/domain/usecases/ucroom"
	"github.com/joaovds/stop-go/internal/shared/errs"
)

type ListRoomUseCaseImpl struct {
	repo repositories.RoomRepository
}

func NewListRoomUseCaseImpl(repo repositories.RoomRepository) *ListRoomUseCaseImpl {
	return &ListRoomUseCaseImpl{repo: repo}
}

// ----- ... -----

func (uc *ListRoomUseCaseImpl) Execute() ([]*ucroom.ListRoomOutput, *errs.Error) {
	rooms, err := uc.repo.List()
	if err != nil {
		return nil, err
	}

	return ucroom.NewListRoomOutputList(rooms), nil
}
