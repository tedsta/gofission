package core

import (
	"testing"
)

type TestEvent struct {
	myNum int
	ch    chan bool
}

func (t *TestEvent) Type() EventType {
	return 0
}

type TestEventHandler struct {
}

func (t *TestEventHandler) HandleEvent(event Event) {
	if test, ok := event.(*TestEvent); ok && test.myNum == 42 {
		event.(*TestEvent).ch <- true
	} else {
		event.(*TestEvent).ch <- false
	}
}

func TestEventManager(t *testing.T) {
	ch := make(chan bool)
	eventManager := &EventManager{}
	testHandler := &TestEventHandler{}
	eventManager.AddHandler(0, testHandler)
	go eventManager.FireEvent(&TestEvent{42, ch})
	if !<-ch {
		t.Fail()
	}
}
