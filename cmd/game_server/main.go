package main

import (
	"fmt"

	"github.com/joaovds/stop-go/internal/application/usecases/roomcasesimpl"
	"github.com/joaovds/stop-go/internal/infra/database/memory/memrepos"
)

func main() {
	room_repo := memrepos.NewRoomRepository()
	list_room_uc := roomcasesimpl.NewListRoomUseCaseImpl(room_repo)

	rooms, err := list_room_uc.Execute()
	if err != nil {
		panic(err)
	}

	for _, room := range rooms {
		fmt.Printf("Room: %v\n", room)
	}
}
