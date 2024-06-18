package entities

type Room struct {
	ID   int
	Name string
	Code string
}

func NewRoom(name string) *Room {
	return &Room{
		Name: name,
	}
}

func MakeRoom(id int, name, code string) *Room {
	room := NewRoom(name)
	room.ID = id
	room.Code = code
	return room
}
