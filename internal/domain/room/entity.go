package room

import (
	"crypto/rand"
	"math/big"
	"time"

	"github.com/joaovds/stop-go/internal/domain/player"
	"github.com/joaovds/stop-go/pkg/errs"
)

type Room struct {
	ID           string
	Name         string
	Code         string
	Players      map[string]*player.Player
	TotalPlayers int
	HostID       string
	MaxPlayers   int
	MinPlayers   int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewRoom(id, name string, maxPlayers, minPlayers int, createdBy *player.Player) (*Room, *errs.Error) {
	room := &Room{
		ID:           id,
		Name:         name,
		Players:      map[string]*player.Player{createdBy.ID: createdBy},
		HostID:       createdBy.ID,
		MaxPlayers:   maxPlayers,
		MinPlayers:   minPlayers,
		TotalPlayers: 1,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := room.generateCode(); err != nil {
		return nil, err
	}

	if err := room.Validate(); err != nil {
		return nil, err
	}

	return room, nil
}

var (
	ErrInvalidRole                     = errs.NewError("invalid role").SetStatus(400)
	ErrRoomFull                        = errs.NewError("room is full").SetStatus(400)
	ErrNameAlreadyExists               = errs.NewError("name already exists").SetStatus(409)
	ErrInvalidMaxPlayers               = errs.NewError("max_players must be greater than 1").SetStatus(400)
	ErrInvalidMinPlayers               = errs.NewError("min_players must be greater than 2").SetStatus(400)
	ErrMinPlayersGreaterThanMaxPlayers = errs.NewError("min_players must be less than max_players").SetStatus(400)
	ErrGenerateCode                    = errs.NewError("error generating code").SetStatus(500)
)

func (r *Room) Validate() *errs.Error {
	if r.MaxPlayers <= 0 {
		return ErrInvalidMaxPlayers
	}
	if r.MinPlayers <= 0 {
		return ErrInvalidMinPlayers
	}
	if r.MinPlayers > r.MaxPlayers {
		return ErrMinPlayersGreaterThanMaxPlayers
	}

	return nil
}

func (r *Room) AddPlayer(p *player.Player) *errs.Error {
	if r.IsFull() {
		return ErrRoomFull
	}

	r.Players[p.ID] = p
	r.TotalPlayers++
	return nil
}

func (r *Room) generateCode() *errs.Error {
	code := make([]byte, 6)
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	for i := range code {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return ErrGenerateCode
		}
		code[i] = charset[num.Int64()]
	}

	r.Code = string(code)

	return nil
}

// ----- ... -----

func (r *Room) IsFull() bool {
	return r.TotalPlayers >= r.MaxPlayers
}
