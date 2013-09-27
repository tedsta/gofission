package event

import (
	"testing"
)

type TestEvent struct {
	myNum int
	ch    chan bool
}

func (t *TestEvent) Type() Type {
	return 0
}

type TestListener struct {
}

func (t *TestListener) HandleEvent(event Event) {
	if test, ok := event.(*TestEvent); ok && test.myNum == 42 {
		event.(*TestEvent).ch <- true
	} else {
		event.(*TestEvent).ch <- false
	}
}

func TestEventManager(t *testing.T) {
	eventManager := &Manager{}

	testListener := &TestListener{}
	eventManager.AddHandler(0, testListener)

	ch := make(chan bool)
	go eventManager.FireEvent(&TestEvent{42, ch})
	if !<-ch {
		t.Fail()
	}
}
