package room

type Room struct {
	ID   int
	Name string
}

func NewRoom(name string) *Room {
	return &Room{
		Name: name,
	}
}
