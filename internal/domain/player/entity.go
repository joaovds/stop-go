package player

import "time"

type Player struct {
	ID         int
	Nickname   string
	Created_at time.Time
	Update_at  time.Time
}

func NewPlayer(nickname string) *Player {
	return &Player{
		Nickname:   nickname,
		Created_at: time.Now(),
		Update_at:  time.Now(),
	}
}
