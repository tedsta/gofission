package core

import (
	"testing"
)

type TestEvent struct {
	myNum int
}

func (t *TestEvent) Type() int {
	return 0
}

type TestEventHandler struct {
	ch chan bool
}

func (t *TestEventHandler) HandleEvent(event Event) {
	if test, ok := event.(*TestEvent); ok && test.myNum == 42 {
		t.ch <- true
	} else {
		t.ch <- false
	}
}

func TestEventManager(t *testing.T) {
	eventManager := NewEventManager(1)
	testHandler := &TestEventHandler{make(chan bool)}
	eventManager.AddHandler(0, testHandler)
	eventManager.FireEvent(&TestEvent{42})
	if !<-testHandler.ch {
		t.Fail()
	}
}
