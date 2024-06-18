package repositories

import (
	"github.com/joaovds/stop-go/internal/domain/entities"
	"github.com/joaovds/stop-go/internal/shared/errs"
)

type RoomRepository interface {
	List() ([]entities.Room, *errs.Error)
}
