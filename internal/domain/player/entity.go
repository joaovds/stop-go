package player

import (
	"time"

	"github.com/joaovds/stop-go/pkg/errs"
)

type PlayerRole string

func (r PlayerRole) String() string {
	return string(r)
}

func (r PlayerRole) IsValid() bool {
	switch r {
	case Host, Adm, Member:
		return true
	}
	return false
}

const (
	Host   PlayerRole = "HOST"
	Adm    PlayerRole = "ADMIN"
	Member PlayerRole = "MEMBER"
)

type Player struct {
	ID        string
	Nickname  string
	Role      PlayerRole
	Score     int
	JoinedAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPlayer(id, nickname string, role PlayerRole) (*Player, *errs.Error) {
	player := &Player{
		ID:        id,
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
	ErrPlayerAlreadyExists   = errs.NewError("player already exists").SetStatus(409)
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
