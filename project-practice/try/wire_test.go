package try

import (
	"testing"
)

func TestUsualCombine(t *testing.T) {
	msg := &Message{"How are you?"}
	operation := &Operator{msg, 1}
	e := &Event{operation}
	t.Log(e.Dispatch())
}

func TestWireCombine(t *testing.T) {
	e, _ := InitializeEvent("How are you?", 1)
	t.Log(e.Dispatch())
}
