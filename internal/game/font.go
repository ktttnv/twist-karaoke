package game

import (
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var cyrillicFace *text.GoTextFace

func init() {
	f, err := os.Open("internal/game/assets/Circe-Regular.ttf")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	src, err := text.NewGoTextFaceSource(f)
	if err != nil {
		panic(err)
	}
	cyrillicFace = &text.GoTextFace{
		Source: src,
		Size:   14,
	}
}

func drawText(screen *ebiten.Image, s string, x, y int, size float64, c color.Color) {
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(float64(x), float64(y))
	opts.ColorScale.ScaleWithColor(c)
	face := &text.GoTextFace{
		Source: cyrillicFace.Source,
		Size:   size,
	}
	text.Draw(screen, s, face, opts)
}
