package fsm

import (
	"golang.org/x/mobile/sprite"
	"image"
	"image/color"
	"image/draw"
)

// LoadColorTexture creates a sprite.SubTex from a standard color.Color.
func LoadColorTexture(eng sprite.Engine, c color.Color, width, height int) (sprite.SubTex, error) {
	img := image.Rect(0, 0, width, height)
	m := image.NewRGBA(img)
	draw.Draw(m, m.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	t, err := eng.LoadTexture(m)
	if err != nil {
		return sprite.SubTex{}, err
	}
	return sprite.SubTex{t, img}, nil
}
