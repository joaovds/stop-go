package memdata

import (
	"time"

	"github.com/joaovds/stop-go/internal/domain/player"
)

var Players = make(map[string]*Player)

type Player struct {
	ID        string
	Nickname  string
	Role      string
	Score     int
	JoinedAt  time.Time
	InRoom    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ----- ... -----

func (p *Player) ToDomain() *player.Player {
	return &player.Player{
		ID:        p.ID,
		Nickname:  p.Nickname,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
