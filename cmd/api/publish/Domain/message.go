package domain

type Message struct {
	Description string
	Name        string
	Status      bool
}

func NewMessage(description, name string, status bool) *Message {
	return &Message{
		Description: description,
		Name:        name,
		Status:      status,
	}
}
