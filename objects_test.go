package fsm

import (
	"testing"
)

func TestCollide(t *testing.T) {
	o := &Object{X: 10, Y: 10, Width: 10, Height: 10}
	o1 := &Object{X: 0, Y: 0, Width: 10, Height: 10}
	o2 := &Object{X: 5, Y: 5, Width: 10, Height: 10}

	if c := o.Collide(o1); c {
		t.Errorf("Collide o,o1 %t, want %t", c, false)
	}
	if c := o.Collide(o2); !c {
		t.Errorf("Collide o,o2 %t, want %t", c, true)
	}
	if c := o.Collide(o); !c {
		t.Errorf("Collide o,o %t, want %t", c, true)
	}
}
