package room

type Room struct {
	ID   string
	Name string
}

func NewRoom(name string) *Room {
	return &Room{
		Name: name,
	}
}
