package fission

import (
	"testing"
)

type TestEvent struct {
	myNum int
}

func (this *TestEvent) Type() int {
	return 0
}

type TestEventHandler struct {
	ch chan bool
}

func (this *TestEventHandler) HandleEvent(event Event) {
	if test, ok := event.(*TestEvent); ok && test.myNum == 42 {
		this.ch <- true
	} else {
		this.ch <- false
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
