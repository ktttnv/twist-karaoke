package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/ktttnv/twist-karaoke/internal/audio"
)

type Game struct {
	capture *audio.Capture
}

func New(c *audio.Capture) *Game {
	return &Game{capture: c}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	rawVol := g.capture.RMS()
	smoothVol := g.capture.SmoothedRMS()

	rawVolHeight := float64(rawVol) * 500

	rawX := 50
	rawY := 600 - rawVolHeight
	rawW := 80
	rawH := rawVolHeight

	vector.FillRect(screen, float32(rawX), float32(rawY), float32(rawW), float32(rawH), color.RGBA{255, 100, 100, 255}, false)

	smoothVolHeight := float64(smoothVol) * 500

	smoothX := 150
	smoothY := 600 - smoothVolHeight
	smoothW := 80
	smoothH := smoothVolHeight

	vector.FillRect(screen, float32(smoothX), float32(smoothY), float32(smoothW), float32(smoothH), color.RGBA{0, 255, 0, 255}, false)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Volume raw: %.4f\nVolume smooth: %.4f", rawVol, smoothVol))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}
