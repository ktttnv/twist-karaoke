package game

import "github.com/hajimehoshi/ebiten/v2"

type Screen interface {
	Update() (next Screen, err error)
	Draw(screen *ebiten.Image)
}
