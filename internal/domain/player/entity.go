package player

import (
	"time"

	"github.com/joaovds/stop-go/pkg/errs"
)

type Player struct {
	ID        int
	Nickname  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPlayer(nickname string) (*Player, *errs.Error) {
	player := &Player{
		Nickname:  nickname,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := player.Validate(); err != nil {
		return nil, err
	}

	return player, nil
}

// ----- ... -----

var (
	ErrInvalidNickname = errs.NewError("invalid nickname").SetStatus(400)
)

func (p *Player) Validate() *errs.Error {
	if p.Nickname == "" {
		return ErrInvalidNickname
	}

	return nil
}
