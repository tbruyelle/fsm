package fsm

import (
	"golang.org/x/mobile/f32"
	"golang.org/x/mobile/sprite"
	"golang.org/x/mobile/sprite/clock"
)

type Object struct {
	// Position
	X, Y float32
	// Speed
	Vx, Vy float32
	// Rotation
	Rx, Ry, Angle float32
	// Translation
	Tx, Ty float32
	// Scale
	Sx, Sy, Scale float32
	Width, Height float32
	Sprite        sprite.SubTex
	Action        Action
	Dead          bool
	Time          clock.Time
	// Data contains any relevant information needed about the object
	Data interface{}
}

type Action interface {
	Do(o *Object, t clock.Time)
}

// Node creates a registered node from the Object and append it.
func (o *Object) Node(parent *sprite.Node, eng sprite.Engine) {
	n := &sprite.Node{}
	eng.Register(n)
	parent.AppendChild(n)
	n.Arranger = o
}

func (o *Object) Reset() {
	o.Tx, o.Ty = 0, 0
	o.Sx, o.Sy, o.Scale = 0, 0, 0
	o.Rx, o.Ry, o.Angle = 0, 0, 0
	o.Vx, o.Vy = 0, 0
	o.Time = 0
}

func (o *Object) Stop() {
	o.Vx, o.Vy = 0, 0
}

func (o *Object) Arrange(e sprite.Engine, n *sprite.Node, t clock.Time) {
	if o.Action != nil {
		// Invoke the action
		o.Action.Do(o, t)
	}

	if o.Dead {
		// Do nothing if dead object
		e.SetSubTex(n, sprite.SubTex{})
		return
	}

	// Set the texture
	e.SetSubTex(n, o.Sprite)

	// Compute affine transformations
	mv := &f32.Affine{}
	mv.Identity()

	// Apply rotations and scales
	if o.Angle != 0 && o.Rx == o.Sx && o.Ry == o.Sy {
		// Optim when angle and scale use the same transformation
		mv.Translate(mv, o.Rx+o.Tx, o.Ry+o.Ty)
		mv.Rotate(mv, -o.Angle)
		mv.Scale(mv, o.Scale, o.Scale)
		mv.Translate(mv, -o.Rx-o.Tx, -o.Ry-o.Ty)
	} else {
		if o.Angle != 0 {
			mv.Translate(mv, o.Rx+o.Tx, o.Ry+o.Ty)
			mv.Rotate(mv, -o.Angle)
			mv.Translate(mv, -o.Rx-o.Tx, -o.Ry-o.Ty)
		}
		if o.Sx > 0 || o.Sy > 0 {
			mv.Translate(mv, o.Sx+o.Tx, o.Sy+o.Ty)
			mv.Scale(mv, o.Scale, o.Scale)
			mv.Translate(mv, -o.Sx-o.Tx, -o.Sy-o.Ty)
		}
	}
	// Add the speeds
	o.X += o.Vx
	o.Y += o.Vy
	// Apply translations
	mv.Translate(mv, o.X+o.Tx, o.Y+o.Ty)
	mv.Mul(mv, &f32.Affine{
		{o.Width, 0, 0},
		{0, o.Height, 0},
	})
	e.SetTransform(n, *mv)
}

// Collide performs an AABB collision.
func (o0 *Object) Collide(o *Object) bool {
	if o.Dead {
		return false
	}
	if o.X >= o0.X+o0.Width || o.X+o.Width <= o0.X || o.Y >= o0.Y+o0.Height || o.Y+o.Height <= o0.Y {
		return false
	}
	return true
}
