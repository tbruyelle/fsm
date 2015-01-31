package fsm

import "golang.org/x/mobile/sprite/clock"

type ActionFunc func(o *Object, t clock.Time)

func (a ActionFunc) Do(o *Object, t clock.Time) {
	a(o, t)
}

// Wait pauses the display of the current object
type Wait struct {
	Until clock.Time
	Next  Action
}

func (w Wait) Do(o *Object, t clock.Time) {
	if o.Time == 0 {
		o.Time = t
		o.Dead = true
		return
	}
	if t > o.Time+w.Until {
		// Once the time is elapsed,
		// start the next Action
		o.Time = 0
		o.Dead = false
		o.Action = w.Next
	}
}
