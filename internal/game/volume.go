package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func drawVolumeBar(screen *ebiten.Image, x, y, w, h, vol float64, c color.Color) {
	if vol < 0 {
		vol = 0
	}
	if vol > 1 {
		vol = 1
	}
	barH := vol * h
	vector.FillRect(screen, float32(x), float32(y+h-barH), float32(w), float32(barH), c, false)
}
