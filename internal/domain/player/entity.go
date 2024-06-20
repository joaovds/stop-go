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
	ErrInvalidNickname       = errs.NewError("invalid nickname").SetStatus(400)
	ErrNicknameAlreadyExists = errs.NewError("nickname already exists").SetStatus(409)
	ErrNicknameTooLong       = errs.NewError("nickname too long. Max 20 characters").SetStatus(400)
)

func (p *Player) Validate() *errs.Error {
	if p.Nickname == "" {
		return ErrInvalidNickname
	}

	if len(p.Nickname) > 20 {
		return ErrNicknameTooLong
	}

	return nil
}