package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func drawButton(screen *ebiten.Image, b MenuButton) {
	vector.FillRect(screen, float32(b.rect.Min.X), float32(b.rect.Min.Y), float32(b.rect.Dx()), float32(b.rect.Dy()), color.RGBA{60, 60, 60, 255}, false)
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(float64(b.rect.Min.X+10), float64(b.rect.Min.Y+10))
	opts.ColorScale.ScaleWithColor(color.White)
	text.Draw(screen, b.label, cyrillicFace, opts)
}
