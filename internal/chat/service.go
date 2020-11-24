package chat

import "github.com/maxisantomil/GoLang2020.git/internal/config"

// Message ...
type Message struct {
	ID   int64
	Text string
}

// ChatService ...
type ChatService interface { /*interface*/
	AddMessage(Message) error
	FindByID(int) *Message
	FindAll() []*Message
}

// devuelve algo que es privado
type service struct {
	conf *config.Config
}

// New ...
func New(c *config.Config) (ChatService, error) {
	return service{c}, nil
}

func (s service) AddMessage(m Message) error {
	return nil
}

func (s service) FindByID(int) *Message {
	return nil
}

func (s service) FindAll() []*Message {
	var list []*Message
	list = append(list, &Message{0, "Hello world"}, &Message{1, "Hello maxi"})
	return list
}
