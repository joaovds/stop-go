package memdata

import "time"

type Player struct {
	ID        string
	Nickname  string
	Role      string
	Score     int
	JoinedAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
