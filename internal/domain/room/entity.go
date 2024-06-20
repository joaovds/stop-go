package room

type Room struct {
	ID   string
	Name string
}

func NewRoom(id, name string) *Room {
	return &Room{
		ID:   id,
		Name: name,
	}
}
