package entities

type Player struct {
	ID       int
	Nickname string
}

func NewPlayer(nickname string) *Player {
	return &Player{Nickname: nickname}
}

func MakePlayer(id int, nickname string) *Player {
	player := NewPlayer(nickname)
	player.ID = id
	return player
}
