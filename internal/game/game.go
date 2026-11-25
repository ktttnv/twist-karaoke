package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/ktttnv/twist-karaoke/internal/audio"
)

type Game struct {
	capture *audio.Capture
	current Screen
}

func New(c *audio.Capture) *Game {
	g := &Game{capture: c}
	g.current = NewMenuScreen(g)
	return g
}

func (g *Game) Update() error {
	if g.current == nil {
		return nil
	}
	next, err := g.current.Update()
	if err != nil {
		return err
	}
	if next != nil {
		g.current = next
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.current == nil {
		return
	}

	g.current.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}
