package fsm

import (
	"image"
	"image/color"
	"image/draw"
	"log"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/sprite"
)

const (
	texBg = iota
)

// MustLoadTexture loads a texture from a file or panics.
func MustLoadTexture(eng sprite.Engine, file string) sprite.Texture {
	f, err := app.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	t, err := eng.LoadTexture(img)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

// SubTex returns a new sprite.SubTex cut from a sprite.Texture.
func SubTex(t sprite.Texture, x1, y1, x2, y2 int) sprite.SubTex {
	return sprite.SubTex{t, image.Rect(x1, y1, x2, y2)}
}

// MustLoadColorTexture creates a sprite.SubTex from a standard color.Color or panics.
func MustLoadColorTexture(eng sprite.Engine, c color.Color, width, height int) sprite.SubTex {
	img := image.Rect(0, 0, width, height)
	m := image.NewRGBA(img)
	draw.Draw(m, m.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	t, err := eng.LoadTexture(m)
	if err != nil {
		log.Fatal(err)
	}
	return sprite.SubTex{t, img}
}
