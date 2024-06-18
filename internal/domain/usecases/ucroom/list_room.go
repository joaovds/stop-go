package ucroom

import (
	"github.com/joaovds/stop-go/internal/domain/entities"
	"github.com/joaovds/stop-go/internal/shared/errs"
)

type ListRoomOutput struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

func NewListRoomOutput(room *entities.Room) *ListRoomOutput {
	return &ListRoomOutput{
		ID:   room.ID,
		Name: room.Name,
		Code: room.Code,
	}
}

func NewListRoomOutputList(rooms []*entities.Room) []*ListRoomOutput {
	var list []*ListRoomOutput
	for _, room := range rooms {
		if room == nil {
			continue
		}

		list = append(list, NewListRoomOutput(room))
	}
	return list
}

type ListRoomUseCase interface {
	Execute() ([]*ListRoomOutput, *errs.Error)
}
