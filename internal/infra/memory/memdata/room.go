package memdata

import "time"

var Rooms = make(map[string]*Room)

type Room struct {
	ID        string
	Name      string
	Code      string
	Max       int
	Min       int
	HostID    string
	Password  string
	IsPrivate bool
	Players   map[string]*Player
	CreatedAt time.Time
	UpdatedAt time.Time
}
