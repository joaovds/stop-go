package sqliterepos

import (
	"context"
	"fmt"

	"github.com/joaovds/stop-go/internal/domain/room"
	"github.com/joaovds/stop-go/internal/infra/sqlite"
	"github.com/joaovds/stop-go/pkg/errs"
)

type RoomRepository struct {
	db *sqlite.DB
}

func NewRoomRepository(db *sqlite.DB) *RoomRepository {
	return &RoomRepository{
		db: db,
	}
}

// ----- ... -----

func (r *RoomRepository) FindAll(ctx context.Context) ([]*room.Room, *errs.Error) {
	return []*room.Room{
		{
			ID:   "1",
			Name: "Room 1",
		},
	}, nil
}

// ----- ... -----

func (r *RoomRepository) NameExists(ctx context.Context, name string) (bool, *errs.Error) {
	var exists bool
	err := r.db.QueryRow(nameExistsQuery, name).Scan(&exists)
	if err != nil {
		return false, errs.NewError("error checking if name exists")
	}

	return exists, nil
}

// ----- ... -----

func (r *RoomRepository) Create(ctx context.Context, roomParams *room.Room) *errs.Error {
	tx, err := r.db.Begin()
	if err != nil {
		return errs.NewError("error creating room")
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	_, err = tx.Exec(createRoomQuery, roomParams.ID, roomParams.Name, roomParams.Code, roomParams.MaxPlayers, roomParams.MinPlayers, roomParams.CreatedAt, roomParams.UpdatedAt)
	if err != nil {
		fmt.Println("1 exe", err)
		return errs.NewError("error creating room")
	}

	_, err = tx.Exec(createRoomPlayerQuery, roomParams.ID, roomParams.HostID, room.Host, roomParams.CreatedAt, roomParams.UpdatedAt)
	if err != nil {
		fmt.Println("2 exe", err)
		return errs.NewError("error creating room player")
	}

	err = tx.Commit()
	if err != nil {
		return errs.NewError("error creating room")
	}

	return nil
}
