package fsm

import "golang.org/x/mobile/sprite/clock"

type ActionFunc func(o *Object, t clock.Time)

func (a ActionFunc) Do(o *Object, t clock.Time) {
	a(o, t)
}

// Wait pauses the display of the current object
type Wait struct {
	// Next contains the next action.
	Next Action
	// Time before Next is triggered.
	Until clock.Time
	// DeadDuring represents the state of the object during wait.
	DeadDuring bool
	// DeadAfter is affected to the object at the end of the wait.
	DeadAfter bool
}

func (w Wait) Do(o *Object, t clock.Time) {
	if o.Time == 0 {
		o.Time = t
		o.Dead = w.DeadDuring
		return
	}
	if t > o.Time+w.Until {
		// Once the time is elapsed,
		// start the next Action
		o.Time = 0
		o.Dead = w.DeadAfter
		o.Action = w.Next
	}
}
