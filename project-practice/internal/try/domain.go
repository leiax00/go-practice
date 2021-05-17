package try

import (
	"fmt"
	"github.com/google/wire"
)

type Message struct {
	Msg string
}

func NewMessage(msg string) *Message {
	return &Message{msg}
}

type Operator struct {
	Msg  *Message
	Code int
}

func NewOperator(msg *Message, code int) *Operator {
	return &Operator{msg, code}
}

type Event struct {
	Operator *Operator
}

func NewEvent(operator *Operator) *Event {
	return &Event{operator}
}

func (e *Event) Dispatch() string {
	return fmt.Sprintf("code: %d, message: %v", e.Operator.Code, e.Operator.Msg.Msg)
}

var EventSet = wire.NewSet(NewOperator, NewEvent, NewMessage)
