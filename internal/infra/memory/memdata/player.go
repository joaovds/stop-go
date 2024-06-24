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
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPlayer(p *player.Player) *Player {
	return &Player{
		ID:        p.ID,
		Nickname:  p.Nickname,
		Role:      string(p.Role),
		Score:     p.Score,
		JoinedAt:  p.JoinedAt,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (p *Player) ToDomain() *player.Player {
	return &player.Player{
		ID:        p.ID,
		Nickname:  p.Nickname,
		Role:      player.PlayerRole(p.Role),
		Score:     p.Score,
		JoinedAt:  p.JoinedAt,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}
